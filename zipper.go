package main

import (
	"archive/zip"
	"bufio"
	_ "bytes"
	"fmt"
	"io"
	"os"
)

var zipname string = "list"

func zipper(hashes, domains <-chan string, listDir string) {

	flatFileNames := writeFlatFile(hashes, domains, listDir)
	err := createZipFile(flatFileNames, listDir)
	if err != nil {
		panic(err)
	}
	//remove flat file to save a little space
	for _, filename := range flatFileNames {
		os.Remove(filename)
	}

}

func unzipper() {

}

func createZipFile(flatFileNames [2]string, listDir string) error {
	filename := fmt.Sprintf("%v%c%v.zip", listDir, os.PathSeparator, zipname)
	newfile, err := os.Create(filename)
	newfile.Chmod(os.ModePerm)
	if err != nil {
		return err
	}
	defer newfile.Close()

	zipWriter := zip.NewWriter(newfile)
	defer zipWriter.Close()

	for _, flatFileName := range flatFileNames {

		zipfile, err := os.Open(flatFileName)
		if err != nil {
			return err
		}
		defer zipfile.Close()

		// Get the file information
		info, err := zipfile.Stat()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Change to deflate to gain better compression
		// see http://golang.org/pkg/archive/zip/#pkg-constants
		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, zipfile)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeFlatFile(hashes, domains <-chan string, listDir string) (filenames [2]string) {
	// I know this is stupid but write everything to a real file
	hashFileName := fmt.Sprintf("%v%c%v.md5", listDir, os.PathSeparator, zipname)
	domainFileName := fmt.Sprintf("%v%c%v.domains", listDir, os.PathSeparator, zipname)
	filenames = [2]string{hashFileName, domainFileName}

	// create flat md5 file.
	file, _ := os.Create(hashFileName)
	chmodErr := file.Chmod(os.ModePerm)
	if chmodErr != nil {
		panic(chmodErr)
	}
	writer := bufio.NewWriter(file)

	for hashedLine := range hashes {
		//fmt.Printf(".")
		writer.Write([]byte(hashedLine + "\n"))
	}

	writer.Flush()
	file.Close()

	//create flat domains file
	file, _ = os.Create(domainFileName)
	chmodErr = file.Chmod(os.ModePerm)
	if chmodErr != nil {
		panic(chmodErr)
	}
	writer = bufio.NewWriter(file)

	for hashedLine := range domains {
		//fmt.Printf(".")
		writer.Write([]byte(hashedLine + "\n"))
	}

	writer.Flush()
	file.Close()

	return
}
