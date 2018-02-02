package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	btgo "github.com/4d55397500/btgo"
	"io/ioutil"
	"os"
	"strconv"
)

var hashKeySize = 4

type ind struct {
	storage btgo.Node
	name    string
}

/*
* constructor for index type
* @param name string
* return index ind
 */
func newIndex(name string) (index *ind) {
	index = new(ind)
	index.name = name
	index.storage = btgo.Node{}
	return
}

/*
index method to add a key value pair to the storage
@param key string
@param value string
*/
func (index *ind) add(value string) {
	ivalue64, _ := strconv.ParseInt(value, 16, 64)
	ivalue := int(ivalue64)
	index.storage.Insert(ivalue)
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
		index.storage = btgo.Node{}
	}

	return err
}

/*
* checks if hash is already in storage returns bool
 */
func (index *ind) contains(value string) bool {
	ivalue64, _ := strconv.ParseInt(value, 16, 64)
	ivalue := int(ivalue64)
	fmt.Println(ivalue)
	return index.storage.Lookup(ivalue)
}

/*
* function so index type implements the io.Writer interface
 */
// func (index *ind) read() (hashes chan string) {
// 	hashes = make(chan string, 1000)
// 	go func() {
// 		defer close(hashes)
// 		for _, arr := range index.storage {
// 			for _, hashSlice := range arr {
// 				hashes <- hashSlice
// 			}
// 		}
// 	}()
// 	return
// }
