package main

import (
	"bufio"
	_ "fmt"
	_ "io"
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

func readUploadByLine(uploadName string, lineChan *chan string) {
	upload, err := os.Open(uploadName)
	defer upload.Close()
	if err != nil {
		panic(err)
	}

	// reader := bufio.NewReader(upload)

	// for {
	//     line, err := reader.ReadString('\n')
	//     if err != nil {
	//             if err == io.EOF {
	//                 break
	//             } else {
	//                 fmt.Println(err)
	//                 os.Exit(1)
	//             }
	//         }
	//     *lineChan <- line
	// }

	scanner := bufio.NewScanner(upload)
	scanner.Buffer([]byte{}, 50)

	for scanner.Scan() {
		*lineChan <- scanner.Text()
	}

	defer close(*lineChan)
}
