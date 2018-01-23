package main

import(
    "fmt"
    "time"
)

func main(){
    var listDir, uploadName string
    var recs, dupes, newRecs int
    var newHashes []string

    // get command line arguments
    uploadName, listDir = getArgs()

    // create index
    index := ind{
        name:listDir+"index",
        storage:map[string][]string{},
    }
    index.open(listDir)

    // get upload file pointers 
    // close file once the function is done
    upload := openUpload(uploadName)
    defer upload.Close()
    
    // read through the file an get info
    timeScan := time.Now()
    recs, newRecs, dupes, newHashes = scanUpload( index, upload )
    timeAfterScan := time.Now()
    totalTimeScan := timeAfterScan.Sub( timeScan )
    fmt.Println( "Time Scan")
    fmt.Println( totalTimeScan )

    // save the index
    timeWrite := time.Now()
    index.write()
    timeAfterWrite := time.Now()
    totalTimeWrite := timeAfterWrite.Sub( timeWrite );
    fmt.Println( "Time write index")
    fmt.Println( totalTimeWrite )

    if newRecs > 0 {
        timeWriteHashes := time.Now()
        writeNewHashes( listDir, newHashes )
        timeAfterWriteHashes := time.Now()
        totalTimeHashes := timeWriteHashes.Sub( timeAfterWriteHashes )
        fmt.Println( "Time write hashes")
        fmt.Println(totalTimeHashes)
    }

    fmt.Printf( "Records: %v New: %v Dupes: %v\n", recs, newRecs, dupes )
    //fmt.Println( index.storage )    

}