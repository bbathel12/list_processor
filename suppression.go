package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

var wg sync.WaitGroup

//Consts
const usage string = "Usage: suppression [-r -h -b -w] -if=<inputFile> -of=<outputDirectory>"

//Channels bufffered seems faster
var buffersize int
var lineChan chan string
var hashedLineChan chan string
var newHashChan chan string
var scanDone chan bool = make(chan bool) // not buffered to keep main routine from finishing

//ints
var recs, newRecs, dupes, workers int

//strings
var listDir, uploadDirectory string
var indexName string = "/GoIndex"

//bools make reindexing work later
var reindex bool
var profile bool

//index
var index *ind

type output struct {
	Totaltime     string `json:"total_time"`
	New_records   int    `json:"new_records"`
	Duplicates    int    `json:"dupes"`
	Total_records int    `json:"records"`
}

func init() {
	//do stuff before main

}

func main() {

	uploadDirectory, listDir, reindex, profile, buffersize, workers = getArgs()
	lineChan = make(chan string, buffersize)
	hashedLineChan = make(chan string, buffersize)
	newHashChan = make(chan string, buffersize)
	recs, newRecs, dupes = 0, 0, 0
	wg.Add(workers)

	start := time.Now()

	if reindex {
		reIndex(uploadDirectory)
		os.Exit(0)
	}

	if profile {
		//Profiling
		f, err := os.Create("./profile/cpu")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}

		defer f.Close()
		defer pprof.StopCPUProfile()
	}

	index = newIndex(listDir + indexName)
	index.open()

	go readUploadDirectory(uploadDirectory, &lineChan)

	// spawn workers for forcing Md5 on lineChan
	for i := 0; i < workers; i++ {
		go loopForceMd5(&lineChan, &hashedLineChan)
	}

	go checkIndex(&recs, &newRecs, &dupes, index, &hashedLineChan, &newHashChan)
	go writeNewHashes(listDir, &newHashChan, &scanDone)

	wg.Wait()             // wait for all forceMd5 routines to
	close(hashedLineChan) // close hashedLineChan which allows checkIndex to finish

	<-scanDone

	index.writeIndexFile()

	// create zip file for download named list.zip
	zipper(index.read(), listDir)

	end := time.Now()
	total := end.Sub(start)

	var stats *output = &output{
		Totaltime:     fmt.Sprintf("%v", total),
		New_records:   newRecs,
		Duplicates:    dupes,
		Total_records: recs,
	}
	stats_bytes, _ := json.Marshal(stats)
	fmt.Println(string(stats_bytes))

	if profile {
		//memory profiling
		fmem, err := os.Create("./profile/memory")
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(fmem); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		fmem.Close()
	}

}
