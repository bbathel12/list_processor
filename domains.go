package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	_ "strings"
)

var scrubDomains map[string]bool
var blank_byte = []byte{}
var scrubReplRegex, _ = regexp.Compile("^[*]?@")

func init() {
	scrubDomains = readScrubDomains()
}

func readScrubDomains() (scrubDomains map[string]bool) {
	config, err := os.Open("./domainsconfig.txt")
	scrubDomains = map[string]bool{}
	if err == nil {

		scanner := bufio.NewScanner(config)
		for scanner.Scan() {
			if _, ok := scrubDomains[scanner.Text()]; !ok {
				scrubDomains[scanner.Text()] = true
			}
		}

	}
	return

}

func inScrubDomains(value string) bool {
	var piece []byte
	piece = scrubReplRegex.ReplaceAll([]byte(value), blank_byte)

	if _, ok := scrubDomains[string(piece)]; !ok {
		return false
	}

	return true
}

func domainLoop(domainChan, newDomainChan *chan string, index *ind) {

	var domain string
	var added bool

	for domain = range *domainChan {
		if !inScrubDomains(domain) {
			if added = index.checkAndAddDomain(domain, &newRecs); added {
				*newDomainChan <- domain
			}
		}

	}

	defer close(*newDomainChan)
}

/*
* writes all new Domains to the a timestamped file in listDir
* @param listDir string: directory of the unique list
* @param newHashes []string: an array containing all hashes to write
 */
func writeNewDomains(listDir string, newDomainChan *chan string, domainScanDone *chan bool) {
	var outFile *os.File
	var writer *bufio.Writer

	for v := range *newDomainChan {
		// create file if not created
		if outFile == nil {
			outFile = openWriteFile(listDir, ".domains")
			writer = bufio.NewWriter(outFile)
		}

		line := fmt.Sprintf("%v\n", v)
		writer.WriteString(line)
		writer.Flush()

	}

	defer outFile.Close()
	defer close(*domainScanDone)
}
