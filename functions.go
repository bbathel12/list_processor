package main

import (
    "bufio"
    "crypto/md5"
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    _"regexp"
    "strings"
    "time"
)

/*
* takes a name of new upload file and opens
* @param filename string 
* @return upload *os.File
*/
func openUpload( filename string ) ( upload *os.File ){
    upload, err := os.Open(filename);
    if err != nil {
        panic( err )
    }
    return 
}

/*
* takes list directory, opens and loads index into memory
* @param listDir string: the directory of the list being updated
* @return outFile *os.File
*/
func openWriteFile(listDir string) (outFile *os.File) {
    t := time.Now()
    filename := fmt.Sprintf( "%s/%d.md5", listDir, t.Unix() )
    outFile , err := os.Create(filename)
    if err != nil {
        panic( err )
    }
    return 
}

/*
* takes string to determin if it is Md5 encoded or not
* @param line string
* @return match bool
*/
func isMd5( line string ) (match bool) {
    match = md5Regex.MatchString( line )
    return 
}

/*
* get command line arguments or die with usage help
*/
func getArgs() ( uploadName, listDir string, runIndexer bool ){
    // get command line arguments
    var reindex *bool
    reindex = flag.Bool("r", false, "re-index list directory")
    inFile  := flag.String("if","", "specify input file with full path")
    outDir := flag.String("of","", "specify output list directory")
    flag.Parse();

    if *inFile == "" || *outDir == "" {
        fmt.Println(usage)
        os.Exit(0);
    }else{
        uploadName = *inFile
        listDir = *outDir
        runIndexer = *reindex
    }
    return 
}


/* 
* reads uploadeName file and goes through each line checking agains the index for
* new records and duplicates
* @param index ind
* @param uploadName string: name of uploaded file 
* @return recs, newRecs, dupes int: number of total records, new records, and duplicate records
* @return newHashes []string: an array of all new hashes
*/
func scanUpload(index ind, uploadName string) {
    // scan through the file and get stuff
    var trimmed string
    var recs, dupes, newRecs int

    uploadBytes, err := ioutil.ReadFile(uploadName)
    uploadString := string(uploadBytes)
    if err != nil{
        panic( err )
    }
    splitUpload := strings.Split(uploadString, "\n")
    for _, line := range splitUpload {
        if line == "" { continue }
        trimmed = strings.TrimSpace(line);
        if  isMd5( trimmed ) {
            if index.contains( trimmed ){
                recs++
                dupes++
            }else{
                newHashChan <- trimmed
                //newHashes = append( newHashes, trimmed )
                index.add( trimmed )
                recs++
                newRecs++
            }
        }else{
            //trimmed = fmt.Sprintf("%v",trimmed)
            bytes := []byte(trimmed)
            hashedBytes := md5.Sum( bytes )
            hashedTrimmed := fmt.Sprintf( "%x", hashedBytes )
            if index.contains( hashedTrimmed ){
                recs++
                dupes++
            }else{
                newHashChan <- hashedTrimmed
                //newHashes = append( newHashes, hashedTrimmed )
                index.add( hashedTrimmed )
                recs++
                newRecs++
            }
        }
    }
    close( newHashChan )
    fmt.Printf( "Records: %v New: %v Dupes: %v\n", recs, newRecs, dupes )
}

/*
* writes all newHashes to the a timestamped file in listDir
* @param listDir string: directory of the unique list
* @param newHashes []string: an array containing all hashes to write
*/
func writeNewHashes( listDir string ) {
    outFile := openWriteFile(listDir)
    writer := bufio.NewWriter(outFile)
    for {
        v, ok :=  <-newHashChan;
        if( !ok ){
            scanDone <- true
            break
        }
        line := fmt.Sprintf( "%v\n", v )
        writer.WriteString( line )
        writer.Flush()

    }
    
    defer outFile.Close()
}



/*
* goes through all files in directory to build index incase 
* lost or never built before
*/
func reindex(listDir string, index ind ) ind {

    files, err := ioutil.ReadDir(listDir)
    if err != nil{
        panic( err )
    }
    for _, file := range files{
        if name := file.Name(); strings.Contains(name,".md5") || strings.Contains(name, ".txt"){
            // scan through the file and add to index but don't write lists.
            var trimmed string
            fileWithPath := listDir+name
            uploadBytes, err := ioutil.ReadFile(fileWithPath)
            uploadString := string(uploadBytes)
            if err != nil{
                panic( err )
            }
            splitUpload := strings.Split(uploadString, "\n")
            for _, line := range splitUpload {
                if line == "" { continue }
                trimmed = strings.TrimSpace(line);
                if  isMd5( trimmed ) {
                    if !index.contains( trimmed ){
                        index.add( trimmed )
                    }
                }else{
                    bytes := []byte(trimmed)
                    hashedBytes := md5.Sum( bytes )
                    hashedTrimmed := fmt.Sprintf( "%x", hashedBytes )
                    if !index.contains( hashedTrimmed ){
                        index.add( hashedTrimmed )
                    }
                }
            }
        }
    }
    return index
}






