package main

import (
	"fmt"
	"testing"
)

func Test_newIndex(t *testing.T) {
	testName := "testName"
	index := newIndex(testName)
	if index.name != testName {
		t.Error("name not correct")
	}
}

func Test_add(t *testing.T) {
	testIndex := newIndex("GoIndex")
	testKeyValues := map[string]string{
		"oneo": "oneone",
		"twot": "twotwo",
		"thre": "threethree",
	}

	// new to test adding new keys to existing index
	testKeyValues3 := map[string]string{
		"four": "fourfour",
		"five": "fivefive",
		"sixs": "sixsix",
	}
	for _, v := range testKeyValues {
		testIndex.add(v)
	}

	for _, v := range testKeyValues {
		testIndex.add(v)
	}

	if len(testIndex.Storage) > 3 {
		t.Error("added duplicate")
	}

	for _, v := range testKeyValues3 {
		testIndex.add(v)
	}

	for k, _ := range testKeyValues {
		if _, ok := testIndex.Storage[k]; !ok {
			errorString := fmt.Sprintf("key %s not found ", k)
			t.Error(errorString)
		}
	}
	for k, _ := range testKeyValues3 {
		if _, ok := testIndex.Storage[k]; !ok {
			errorString := fmt.Sprintf("key %s not found ", k)
			t.Error(errorString)
		}
	}

}

func Test_writeIndex(t *testing.T) {
	index := ind{
		name: "./tests/index",
		Storage: map[string][]string{
			"one":   []string{"1", "2", "3"},
			"two":   []string{"1", "2", "3"},
			"three": []string{"1", "2", "3"},
		},
	}
	index.writeIndexFile()

}

func Test_openIndex(t *testing.T) {
	index := newIndex("GoIndex")
	err := index.open()
	if err != nil {
		t.Error(err)
	}
	//fmt.Println(index.Storage)
}

func Test_contains(t *testing.T) {
	index := ind{
		name: "testindex",
		Storage: map[string][]string{
			"onet":  []string{"onetwo", "onetlu", "onetru"},
			"two":   []string{"1", "2", "3"},
			"three": []string{"1", "2", "3"},
		},
	}
	v := "onetwo"
	if !index.contains(v) {
		t.Error("INDEX DOESN'T CONTAIN " + v)
	}
	v = "orangeorange"
	if index.contains(v) {
		t.Error("INDEX SHOULDN'T CONTAIN " + v)
	}
}

func Test_checkAndAddDomain(t *testing.T) {
	newRecs := new(int)
	index := ind{
		name: "testIndex",
		Domains: map[string]bool{
			"one":   true,
			"two":   true,
			"three": true,
		},
	}

	if index.checkAndAddDomain("one", newRecs) {
		t.Error("one not found")
	}
	if index.checkAndAddDomain("two", newRecs) {
		t.Error("two not found")
	}
	if !index.checkAndAddDomain("asdfasdfa", newRecs) {
		t.Error("asdfasdfa found and shouldn't be")
	}
	if index.checkAndAddDomain("asdfasdfa", newRecs) {
		t.Error("asdfasdfa not found and should be")
	}

}

func Test_read(t *testing.T) {
	index := ind{
		name: "testindex",
		Storage: map[string][]string{
			"onet":  []string{"onetwo", "onetlu", "onetru"},
			"two":   []string{"1", "2", "3"},
			"three": []string{"1", "2", "3"},
		},
	}
	count := 0
	for _ = range index.read() {
		count++
	}
	if count != 9 {
		t.Error("Count is not correct", count)
	}
}

func Test_readDomains(t *testing.T) {
	index := ind{
		name: "testindex",
		Domains: map[string]bool{
			"onet":  true,
			"two":   true,
			"three": true,
		},
	}
	count := 0
	for _ = range index.readDomains() {
		count++
	}
	if count != 3 {
		t.Error("Count is not correct", count)
	}
}

func Test_openWriteFile(t *testing.T) {
	outFile := openWriteFile("./testCase/", ".txt")
	if outFile == nil {
		t.Error("No File Opened")
	}
}

func Benchmark_index_add(b *testing.B) {
	index := newIndex("GoIndex")
	for i := 0; i < b.N; i++ {
		index.add(string(i))
	}
}

func Benchmark_index_contains(b *testing.B) {
	index := newIndex("GoIndex")
	for i := 0; i < 10000; i++ {
		index.add(string(i))
	}

	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		index.contains(string(j))
	}
}

func Benchmark_open_index(b *testing.B) {
	index := newIndex("testCase/GoIndex")
	for i := 0; i < b.N; i++ {
		index.open()
	}
}
