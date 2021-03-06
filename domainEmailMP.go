package main

import (
	"fmt"
	"regexp"
	"strings"
)

//Regex
var md5Regex, _ = regexp.Compile("^[a-f0-9]{32}$")
var domainRegex, _ = regexp.Compile("^((|\\*)@)[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,})$")
var emailRegex, _ = regexp.Compile("^((|\\*|([a-z0-9_!#$%&\\'+\\/=?^`{|}~-]+\\.?)+)@)[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,})$")

func domainEmailMultiPlex(recs *int, lineChan, domainChan, hashedLineChan *chan string) {
	var line, trimmed string
	for line = range *lineChan {

		// clean the input
		trimmed = strings.TrimSpace(line)
		trimmed = strings.ToLower(trimmed)

		switch {
		case domainRegex.MatchString(trimmed):
			*domainChan <- trimmed
			*recs++
		case emailRegex.MatchString(trimmed):
			trimmed = forceMd5(trimmed)
			*hashedLineChan <- trimmed
			*recs++
		case md5Regex.MatchString(trimmed):
			*hashedLineChan <- trimmed
			*recs++
		default:
			fmt.Println("no match")
			//save as errors maybe?
		}

	}

	defer close(*domainChan)
	defer close(*hashedLineChan)

}
