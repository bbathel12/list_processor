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
	
	if len(value) <= hashKeySize {
		halfhash = value
	} else {
		halfhash = value[:hashKeySize]
	}

	index.storage[halfhash] = append(index.storage[halfhash], value)
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
	if len(line) <= hashKeySize {
		halfhash = line
	} else {
		halfhash = line[:hashKeySize]
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


/*
* function so index type implements the io.Writer interface
*/
func (index *ind) read() ( hashes chan string ){
    hashes = make( chan string , 1000)
    go func(){
        defer close(hashes)
        for _, arr := range index.storage{
            for _, hashSlice := range arr{
                hashes <- hashSlice            
            } 
        }
    }()
    return
} 


