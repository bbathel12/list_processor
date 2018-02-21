package main

import (
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
var lineChan, emailChan, domainChan chan string
var hashedLineChan chan string
var newHashChan, newDomainChan chan string
var scanDone, domainScanDone chan bool = make(chan bool), make(chan bool) // not buffered to keep main routine from finishing

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

func init() {
	//do stuff before main

}

func main() {

	uploadDirectory, listDir, reindex, profile, buffersize, workers = getArgs()
	lineChan = make(chan string, buffersize)
	hashedLineChan = make(chan string, buffersize)
	newHashChan = make(chan string, buffersize)
	newDomainChan = make(chan string, buffersize)
	domainChan = make(chan string, buffersize)
	emailChan = make(chan string, buffersize)
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

	go readUploadDirectory(uploadDirectory, &domainChan, &lineChan)
	go domainEmailMultiPlex(&recs, &lineChan, &domainChan, &hashedLineChan)
	go domainLoop(&domainChan, &newDomainChan, index)
	go checkIndex(&newRecs, &dupes, index, &hashedLineChan, &newHashChan)
	go writeNewHashes(listDir, &newHashChan, &scanDone)
	go writeNewDomains(listDir, &newDomainChan, &domainScanDone)

	<-scanDone
	<-domainScanDone
	fmt.Println(index.Domains)
	index.writeIndexFile()

	// create zip file for download named list.zip
	zipper(index.read(), index.readDomains(), listDir)

	end := time.Now()
	total := end.Sub(start)

	// create new output struct and print json
	output := newOutput(total, newRecs, dupes, recs)
	output.printJson()

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
