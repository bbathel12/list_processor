package main

import (
	"bufio"
    "io/ioutil"
    "os"
	"strings"
)

func readUpload(uploadName string, lineChan *chan string) {
	uploadBytes, err := ioutil.ReadFile(uploadName)
	uploadString := string(uploadBytes)
	if err != nil {
		panic(err)
	}
	splitUpload := strings.Split(uploadString, "\n")
	for _, line := range splitUpload {
		if line == "" {
			continue
		}
		*lineChan <- line
	}
	close(*lineChan)
}


func readUploadByLine(uploadName string, lineChan *chan string){
    upload, err := os.Open(uploadName)
    defer upload.Close()
    if err != nil {
        panic(err)
    }

    reader := bufio.NewScanner(upload)
    
    for reader.Scan() {
        *lineChan <- reader.Text()
    }

    defer close(*lineChan)
}
