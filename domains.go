package main

import (
	"bufio"
	"fmt"
	"os"
)

func domainLoop(domainChan, newDomainChan *chan string, index *ind) {

	var domain string
	var added bool

	for domain = range *domainChan {
		if added = index.checkAndAddDomain(domain, &newRecs); added {
			*newDomainChan <- domain
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
