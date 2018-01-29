package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

/*
* get command line arguments or die with usage help
 */
func getArgs() (uploadName, listDir string, runIndexer, profile bool, buffersize int) {
	// get command line arguments

	reindex := flag.Bool("r", false, "re-index list directory")
	inFile := flag.String("if", "", "specify input file with full path")
	outDir := flag.String("of", "", "specify output list directory")
	runProfiler := flag.Bool("p", false, "run profiler")
	buffer := flag.Int("b", 100, "size of buffered channels")
	flag.Parse()

	if *inFile == "" || *outDir == "" {
		fmt.Println(usage)
		os.Exit(0)
	} else {
		uploadName = *inFile
		listDir = *outDir
		runIndexer = *reindex
		profile = *runProfiler
		buffersize = *buffer
	}
	return
}

/*
* takes list directory, opens and loads index into memory
* @param listDir string: the directory of the list being updated
* @return outFile *os.File
 */
func openWriteFile(listDir string) (outFile *os.File) {
	t := time.Now()
	filename := fmt.Sprintf("%s/%d.md5", listDir, t.Unix())
	outFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return
}

/*
* writes all newHashes to the a timestamped file in listDir
* @param listDir string: directory of the unique list
* @param newHashes []string: an array containing all hashes to write
 */
func writeNewHashes(listDir string, newHashChan *chan string, scanDone *chan bool) {
    var outFile *os.File
    var writer *bufio.Writer
    
	for  v := range *newHashChan {
        // create file if not created
        if( outFile == nil ){
            outFile = openWriteFile(listDir)
            writer = bufio.NewWriter(outFile)
        }

        line := fmt.Sprintf("%v\n", v)
        writer.WriteString(line)
        writer.Flush()
        
	}
    
	defer outFile.Close()
    defer close(*scanDone)
}
