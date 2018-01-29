package main

import(
	"io/ioutil"
	"strings"
)


func readUpload(  uploadName string, lineChan *chan string ){
	uploadBytes, err := ioutil.ReadFile(uploadName)
    uploadString := string(uploadBytes)
    if err != nil{
        panic( err )
    }
    splitUpload := strings.Split(uploadString, "\n")
    for _, line := range splitUpload {
        if line == "" { continue }
        *lineChan <- line
    }
    close(*lineChan)
}