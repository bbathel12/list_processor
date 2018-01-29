package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
)



type ind struct {
	storage map[string][]string
	name    string
}


/*
* constructor for index type
* @param name string
* return index ind
*/
func newIndex(name string) (index *ind){
	index = new(ind)
	index.name = name
	index.storage = map[string][]string{}
	return
}

/*
index method to add a key value pair to the storage
@param key string
@param value string
*/
func (index *ind) add(value string) {
	var halfhash string


	half := len(value) / 2
	if half <= 0 {
		halfhash = value
	} else {
		halfhash = value[:half]
	}

	index.storage[halfhash] = append(index.storage[halfhash], value)
}

/*
* writes encoded gob to file named index in the list directory
* @return nil
*/
func (index *ind) write() {
	if _, err := os.Stat(index.name); os.IsNotExist(err) {
		indexFile, err := os.Create(index.name)
		if err != nil {
			fmt.Println(err)
		}
		indexFile.Close()
	}

	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)

	err := e.Encode(index.storage)
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
		decodeErr := d.Decode(&index.storage)
		if decodeErr != nil {
			panic(decodeErr)
		}
		defer indexFile.Close()
	} else {
		index.storage = map[string][]string{}
	}

	return err
}

/*
* checks if hash is already in storage returns bool
*/
func (index *ind) contains(line string) bool {
	var halfhash string

	half := len(line) / 2
	if half <= 0 {
		halfhash = line
	} else {
		halfhash = line[:half]
	}


	if _, ok := index.storage[halfhash]; ok {
		for _, v := range index.storage[halfhash] {
			if line == v {
				return true
			}
		}
	}
    
	return false
}
