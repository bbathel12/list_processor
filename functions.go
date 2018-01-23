package main

import (
    "fmt"
    "os"
    "bufio"
    "regexp"
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
    filename := fmt.Sprintf( "%s/%d.txt", listDir, t.Unix() )
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
func isMd5(line string) (match bool) {
    match, _ = regexp.MatchString("[A-Fa-f0-9]{32}", line)
    return 
}

/*
* get command line arguments or die with usage help
*/
func getArgs() ( uploadName, listDir string ){
    // get command line arguments
    args := os.Args[1:]
    if len( args ) < 2 {
        fmt.Println("Usage: suppression <inputFile> <outputDirectory>")
        os.Exit(0);
    }else{
        uploadName = args[0]
        listDir = args[1]
    }
    return 
}

func scanUpload(index ind, upload *os.File) (recs, newRecs, dupes int , newHashes []string ){
    // scan through the file and get stuff
    scanner := bufio.NewScanner(upload)
    for scanner.Scan(){
        line := scanner.Text()
        if trimmed := strings.TrimSpace(line); isMd5( trimmed ) {
            if index.contains( trimmed ){
                recs++
                dupes++
            }else{
                newHashes = append( newHashes, trimmed )
                index.add( trimmed )
                recs++
                newRecs++
            }
        }
        if err := scanner.Err(); err != nil{
            panic(err)
        }
    }
    return 
}


func writeNewHashes( listDir string, newHashes []string ) {
    outFile := openWriteFile(listDir)

    for _, v := range newHashes{
        writer := bufio.NewWriter(outFile)
        line := fmt.Sprintf( "%v\n", v )
        writer.WriteString( line )
        writer.Flush()
    }

    defer outFile.Close()
}