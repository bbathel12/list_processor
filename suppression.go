//Suppression Package
//  Checks uploaded files, builds index, removes duplicates and
//  writes new records to directory
//      Usage: suppression [-r] -if=<inputFile> -of=<outputDirectory>:
package main

import(
    "fmt"
    "regexp"
    "time"

)

var newHashChan chan string = make (chan string)
var addNewHashesChan chan string = make (chan string)
var scanDone chan bool = make (chan bool)
var md5Regex, _ = regexp.Compile("^[a-f0-9]{32}$")
var usage = "Usage: suppression [-r] -if=<inputFile> -of=<outputDirectory>"


/*

*/
func main(){
    var listDir, uploadName string
    
    //var newHashes []string
    var runIndexer bool

    


    // get command line arguments
    uploadName, listDir, runIndexer = getArgs()

    // create index
    index := ind{
        name:listDir+"GoIndex",
        storage:map[string][]string{},
    }
    

    if runIndexer{
        timeReindex := time.Now()
        index = reindex( listDir, index )
        timeAfterReindex := time.Now()
        totalTimeReindex := timeAfterReindex.Sub( timeReindex )
        fmt.Println( "Time Reindex")
        fmt.Println( totalTimeReindex )
    }else{
        fmt.Println("Opening Index")
        timeOpenIndex := time.Now()
        index.open(listDir)
        timeAfterOpenIndex := time.Now()
        totalTimeOpenIndex:= timeAfterOpenIndex.Sub( timeOpenIndex )
        fmt.Println( "Time open index")
        fmt.Println(totalTimeOpenIndex)

    }

    

    // get upload file pointers 
    // close file once the function is done
    // upload := openUpload(uploadName)
    // defer upload.Close()

    go writeNewHashes( listDir )
    go addNewHashes( &index)
    
    
    // read through the file an get info
    timeScan := time.Now()
    scanUpload( &index, uploadName )
    timeAfterScan := time.Now()
    totalTimeScan := timeAfterScan.Sub( timeScan )
    fmt.Println( "Time Scan")
    fmt.Println( totalTimeScan )

    
    <-scanDone
    // save the index
    timeWrite := time.Now()
    index.write()
    timeAfterWrite := time.Now()
    totalTimeWrite := timeAfterWrite.Sub( timeWrite );
    fmt.Println( "Time write index")
    fmt.Println( totalTimeWrite )

    
   
   
    //fmt.Println( index.storage )    

}