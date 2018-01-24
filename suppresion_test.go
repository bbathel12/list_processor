package main

import (
	_ "bytes"
	_ "encoding/gob"
	"fmt"
	"testing"
)

func Test_openUpload(t *testing.T) {
	filename := "./tests/testupload.txt"
	file := openUpload(filename)
	if file == nil {
		t.Error("File not found")
	}
}

func Test_openIndex(t *testing.T) {
	index := ind{}
	err := index.open("./testing/")
	if err != nil {
		t.Error(err)
	}
	//fmt.Println(index.storage)
}

func Test_openWriteFile(t *testing.T) {

}

func Test_add(t *testing.T) {
	testIndex := ind{
		name:    "testindex",
		storage: map[string][]string{},
	}
	testKeyValues := map[string]string{
		"one":   "oneone",
		"two":   "twotwo",
		"three": "threethree",
	}
    // duplicates to test duplicate rejection
    testKeyValues2 := map[string]string{
        "one": "oneone",
        "two": "twotwo",
        "three": "threethree",
    }
    // new to test adding new keys to existing index
    testKeyValues3 := map[string]string{
        "four": "fourfour",
        "five": "fivefive",
        "six": "sixsix",
    }
	for _, v := range testKeyValues {
		testIndex.add(v)
	}

    for _, v := range testKeyValues2 {
        testIndex.add(v)
    }

    if len(testIndex.storage) > 3 {
        t.Error("added duplicate")
    }

    for _, v := range testKeyValues3 {
        testIndex.add(v)
    }

    
	for k, _ := range testKeyValues {
		if _, ok := testIndex.storage[k]; !ok {
			errorString := fmt.Sprintf("key %s not found ", k)
			t.Error(errorString)
		}
	}
    for k, _ := range testKeyValues3 {
        if _, ok := testIndex.storage[k]; !ok {
            errorString := fmt.Sprintf("key %s not found ", k)
            t.Error(errorString)
        }
    }


}

func Test_writeIndex(t *testing.T) {
	index := ind{
		name: "./tests/index",
		storage: map[string][]string{
			"one":   []string{"1", "2", "3"},
			"two":   []string{"1", "2", "3"},
			"three": []string{"1", "2", "3"},
		},
	}
	index.write()

}

func Test_isMd5(t *testing.T) {
	
	tests := []string{
		"4A8A08F09D37B73795649038408B5F33",
		"8277E0910D750195B448797616E091AD",
		"E1671797C52E15F763380B45E841EC32",
        "brice@gmail.com",
        "gmail.com",
        "*.tld",
	}
	for _, v := range tests {
        hashedTrimmed := forceMd5(v)
		if !md5Regex.MatchString( hashedTrimmed ) {
			t.Error(v + " is not Md5")
		}
	}
}

func Test_contains(t *testing.T) {
	index := ind{
		name: "testindex",
		storage: map[string][]string{
			"one":   []string{"onetwo", "oneblu", "onetru"},
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

func Test_main(t *testing.T) {
	t.Error("")
}
