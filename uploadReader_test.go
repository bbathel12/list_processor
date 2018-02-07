package main

import (
	"testing"
)

func test_readUpload(t *testing.T) {
	lineChan := make(chan string)
	readUpload("tests/list2.md5", &lineChan)
	testLines := [5]string{
		"0CC175B9C0F1B6A831C399E269772661",
		"92EB5FFEE6AE2FEC3AD71C777531578F",
		"4A8A08F09D37B73795649038408B5F33",
		"8277E0910D750195B448797616E091AD",
		"E1671797C52E15F763380B45E841EC32",
	}
	i := 0
	for {

		v, ok := <-lineChan
		if !ok {
			break
		} // if line
		if testLines[i] != v {
			t.Error(v, " Does Not Match ", testLines[i])
		}
	}
}
