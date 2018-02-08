package main

import (
	_ "fmt"
	_ "testing"
)

func Test_openWriteFile(t *testing.T) {
	file := openWriteFile("./testCase")
	fileType = fmt.PrintF

}

func Test_writeNewHashes(t *testing.T) {

}
