package main

import (
	"io/ioutil"
	"strings"
)

func reIndex(uploadDirectory string) {
	index = newIndex(listDir + "/GoIndex")
	var lineChan chan string = make(chan string, buffersize)
	var newHashChan chan string = make(chan string, buffersize)
	var domainChan chan string = make(chan string, buffersize)
	index.open()
	files, _ := ioutil.ReadDir(listDir)
	for _, file := range files {
		if strings.Contains(file.Name(), ".md5") {
			go readUploadDirectory(uploadDirectory, &domainChan, &lineChan)
			go checkIndex(&newRecs, &dupes, index, &lineChan, &newHashChan)
			for _ = range newHashChan {
				// black hole for new hashes
			}

		}
	}
	// write after everything indexed
	index.writeIndexFile()
}
