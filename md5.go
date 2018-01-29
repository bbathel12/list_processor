package main

import(
	"crypto/md5"
	"fmt"
	"strings"
)


/*
* takes string trims and lowercases it, converts to md5 if not md5
* @param line string
* @return match bool
*/
func forceMd5( lineChan, hashedLineChan *chan string ) {

	for{

		line, ok := <-*lineChan
		if !ok { break; } // if line chan closed break out;

		hashedTrimmed := strings.TrimSpace(line)
	    hashedTrimmed = strings.ToLower(line)

	    if !md5Regex.MatchString( hashedTrimmed ){
	        bytes := []byte(hashedTrimmed)
	        hashedBytes := md5.Sum( bytes )
	        hashedTrimmed = fmt.Sprintf( "%x", hashedBytes )
	    }

	    *hashedLineChan <- hashedTrimmed
	}
	close( *hashedLineChan )

}