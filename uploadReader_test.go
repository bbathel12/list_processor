package main

import (
	_ "fmt"
	"sync"
	"testing"
)

func Test_readUpload(t *testing.T) {
	lineChan := make(chan string, 100)
	var readerGroup sync.WaitGroup

	readerGroup.Add(1)
	go readUpload("tests/list2.md5", &lineChan, &readerGroup)
	readerGroup.Wait()
	close(lineChan)

	testLines := [5]string{
		"0CC175B9C0F1B6A831C399E269772661",
		"92EB5FFEE6AE2FEC3AD71C777531578F",
		"4A8A08F09D37B73795649038408B5F33",
		"8277E0910D750195B448797616E091AD",
		"E1671797C52E15F763380B45E841EC32",
	}
	i := 0

	for v := range lineChan {
		if testLines[i] != v {
			t.Error(v, " Does Not Match ", testLines[i])
		}
		i++
	}

}

func Test_readUploadByLine(t *testing.T) {
	lineChan := make(chan string, 100)
	var readerGroup sync.WaitGroup

	readerGroup.Add(1)
	go readUpload("tests/list2.md5", &lineChan, &readerGroup)
	readerGroup.Wait()
	close(lineChan)

	testLines := [5]string{
		"0CC175B9C0F1B6A831C399E269772661",
		"92EB5FFEE6AE2FEC3AD71C777531578F",
		"4A8A08F09D37B73795649038408B5F33",
		"8277E0910D750195B448797616E091AD",
		"E1671797C52E15F763380B45E841EC32",
	}
	i := 0
	for v := range lineChan {
		if testLines[i] != v {
			t.Error(v, " Does Not Match ", testLines[i])
		}
		i++
	}

	readerGroup.Wait()
}

//
//func Benchmark_readUpload(b *testing.B) {
//
//	for i := 0; i < b.N; i++ {
//
//		var readerGroup sync.WaitGroup
//		readerGroup.Add(1)
//		lineChan := make(chan string, 100)
//		done := make(chan bool, 1)
//		// I think this is just to empty out lineChan
//		go func(lineChan *chan string) {
//			for _ = range *lineChan {
//
//			}
//			readerGroup.Wait()
//			done <- true
//		}(&lineChan)
//
//		readUpload("tests/uploads/supplist.txt", &lineChan, &readerGroup)
//		<-done
//		close(lineChan)
//
//	}
//}
//
//func Benchmark_readUploadByLine(b *testing.B) {
//	var readerGroup sync.WaitGroup
//	readerGroup.Add(1)
//
//	for i := 0; i < b.N; i++ {
//		lineChan := make(chan string, 100)
//		done := make(chan bool, 1)
//		go func(lineChan *chan string) {
//			for _ = range *lineChan {
//
//			}
//			done <- true
//		}(&lineChan)
//		readUploadByLine("tests/uploads/supplist.txt", &lineChan, &readerGroup)
//		<-done
//		readerGroup.Wait()
//		close(lineChan)
//	}
//}
