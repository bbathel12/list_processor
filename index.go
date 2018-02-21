package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
)

var hashKeySize = 4

type ind struct {
	Storage map[string][]string
	name    string
	Domains map[string]bool
}

/*
* constructor for index type
* @param name string
* return index ind
 */
func newIndex(name string) (index *ind) {
	index = new(ind)
	index.name = name
	index.Storage = map[string][]string{}
	index.Domains = map[string]bool{}
	return
}

/*
index method to add a key value pair to the Storage
@param key string
@param value string
*/
func (index *ind) add(value string) {
	var halfhash string

	if len(value) <= hashKeySize {
		halfhash = value
	} else {
		halfhash = value[:hashKeySize]
	}

	index.Storage[halfhash] = append(index.Storage[halfhash], value)
}

/*
* writes encoded gob to file named index in the list directory
* @return nil
 */
func (index *ind) writeIndexFile() {
	if _, err := os.Stat(index.name); os.IsNotExist(err) {
		indexFile, err := os.Create(index.name)
		if err != nil {
			fmt.Println(err)
		}
		indexFile.Close()
	}

	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)

	err := e.Encode(index)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(index.name, b.Bytes(), 7777)
}

/*
* takes list directory, opens and loads index into memory
* @param listDir string: the directory of the list being updated
* @return index ind
* @return err error
 */
func (index *ind) open() (err error) {
	filename := index.name
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		indexFile, _ := os.Open(filename)
		d := gob.NewDecoder(indexFile)
		decodeErr := d.Decode(&index)
		if decodeErr != nil {
			panic(decodeErr)
		}
		defer indexFile.Close()
	} else {
		index.Storage = map[string][]string{}
		index.Domains = map[string]bool{}
	}

	return err
}

/*
* checks if hash is already in Storage returns bool
 */
func (index *ind) contains(line string) bool {
	var halfhash string
	if len(line) <= hashKeySize {
		halfhash = line
	} else {
		halfhash = line[:hashKeySize]
	}

	if _, ok := index.Storage[halfhash]; ok {
		for _, v := range index.Storage[halfhash] {
			if line == v {
				return true
			}
		}
	}

	return false
}

func (index *ind) checkAndAddDomain(domain string, newRecs *int) (added bool) {

	added = false

	if _, ok := index.Domains[domain]; !ok {
		index.Domains[domain] = true
		*newRecs++
		added = true
	}

	return

}

/*
* generator that returns a chan of strings that
* are all the hashses
 */
func (index *ind) read() (hashes chan string) {
	hashes = make(chan string, 1000)
	go func() {
		defer close(hashes)
		for _, arr := range index.Storage {
			for _, hashSlice := range arr {
				hashes <- hashSlice
			}
		}
	}()
	return
}

/*
* generator that returns a chan of strings that
* are all the domains
 */
func (index *ind) readDomains() (domains chan string) {
	domains = make(chan string, 1000)
	go func() {
		defer close(domains)
		for domain, _ := range index.Domains {

			domains <- domain

		}
	}()
	return
}
