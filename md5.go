package main

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"strings"
)

//Regex
var md5Regex, _ = regexp.Compile("^[a-f0-9]{32}$")

/*
* takes string trims and lowercases it, converts to md5 if not md5
* @param line string
* @return match bool
 */
func loopForceMd5(lineChan, hashedLineChan *chan string) {

	for {

		line, ok := <-*lineChan
		if !ok {
			break
		} // if line chan closed break out;

		hashedTrimmed := forceMd5(line)

		*hashedLineChan <- hashedTrimmed
	}
	//close( *hashedLineChan )
	defer wg.Done()

}

func forceMd5(line string) (hashedTrimmed string) {
	hashedTrimmed = strings.TrimSpace(line)
	hashedTrimmed = strings.ToLower(hashedTrimmed)

	if !md5Regex.MatchString(hashedTrimmed) {

		bytes := []byte(hashedTrimmed)
		hashedBytes := md5.Sum(bytes)
		hashedTrimmed = fmt.Sprintf("%x", hashedBytes)

	}
	return
}
