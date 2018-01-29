package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"runtime/pprof"
	"time"
)

//Consts
const usage string = "Usage: suppression [-r -h -b] -if=<inputFile> -of=<outputDirectory>"

//Regex
var md5Regex, _ = regexp.Compile("^[a-f0-9]{32}$")

//Channels bufffered seems faster
var buffersize int
var lineChan chan string
var hashedLineChan chan string
var newHashChan chan string
var scanDone chan bool = make(chan bool) // not buffered to keep main routine from finishing

//ints
var recs, newRecs, dupes int

//strings
var listDir, uploadName string

//bools make reindexing work later
//var reindex bool
var profile bool

//index
var index *ind

func init() {
	// get command line arguments
	uploadName, listDir, _, profile, buffersize = getArgs()
	lineChan = make(chan string, buffersize)
	hashedLineChan = make(chan string, buffersize)
	newHashChan = make(chan string, buffersize)

}

func main() {

	start := time.Now()

	if profile {
		//Profiling
		f, err := os.Create("./profile/cpu")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
		defer f.Close()
	}

	recs, newRecs, dupes = 0, 0, 0

	index = newIndex(listDir + "/GoIndex")
	index.open()

	go readUpload(index, uploadName, &lineChan)
	go forceMd5(&lineChan, &hashedLineChan)
	go checkIndex(&recs, &newRecs, &dupes, index, &hashedLineChan, &newHashChan)
	go writeNewHashes(listDir, &newHashChan, &scanDone)

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
