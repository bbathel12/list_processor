package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

var wg sync.WaitGroup

//Consts
const usage string = "Usage: suppression [-r -h -b -w] -if=<inputFile> -of=<outputDirectory>"

//Regex
var md5Regex, _ = regexp.Compile("^[a-f0-9]{32}$")

//Channels bufffered seems faster
var buffersize int
var lineChan chan string
var hashedLineChan chan string
var newHashChan chan string
var scanDone chan bool = make(chan bool) // not buffered to keep main routine from finishing

//ints
var recs, newRecs, dupes, workers int

//strings
var listDir, uploadName string

//bools make reindexing work later
var reindex bool
var profile bool

//index
var index *ind

func init() {
	// get command line arguments
	uploadName, listDir, reindex, profile, buffersize, workers = getArgs()
	lineChan = make(chan string, buffersize)
	hashedLineChan = make(chan string, buffersize)
	newHashChan = make(chan string, buffersize)
	recs, newRecs, dupes = 0, 0, 0
	wg.Add(workers)

}

func main() {

	start := time.Now()

    // just a timer this can be removed later
	go func() {
		i := 0
		for _ = range time.Tick(time.Second) {
			i++
			fmt.Printf("%v\r", i)
		}
	}()

    if reindex {
        reIndex( listDir )
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

	index = newIndex(listDir + "/GoIndex")
	index.open()

	go readUpload(uploadName, &lineChan)

	// spawn workers for forcing Md5 on lineChan
	for i := 0; i < workers; i++ {
		go forceMd5(&lineChan, &hashedLineChan)
	}

	go checkIndex(&recs, &newRecs, &dupes, index, &hashedLineChan, &newHashChan)
	go writeNewHashes(listDir, &newHashChan, &scanDone)

	wg.Wait()             // wait for all forceMd5 routines to
	close(hashedLineChan) // close hashedLineChan which allows checkIndex to finish

	<-scanDone

	index.write()

	end := time.Now()
	total := end.Sub(start)
	fmt.Printf("Time Total: ")
	fmt.Println(total)
	fmt.Printf("Recs: %d New Recs: %d Dupes: %d ", recs, newRecs, dupes)

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
