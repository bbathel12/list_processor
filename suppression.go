package main

import(
    "fmt"
    "time"
)

func main(){
    var listDir, uploadName string
    var recs, dupes, newRecs int
    var newHashes []string
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
        index.open(listDir)
    }

    

    // get upload file pointers 
    // close file once the function is done
    // upload := openUpload(uploadName)
    // defer upload.Close()
    
    // read through the file an get info
    timeScan := time.Now()
    recs, newRecs, dupes, newHashes = scanUpload( index, uploadName )
    timeAfterScan := time.Now()
    totalTimeScan := timeAfterScan.Sub( timeScan )
    fmt.Println( "Time Scan")
    fmt.Println( totalTimeScan )

    // save the index
    timeWrite := time.Now()
    go index.write()
    timeAfterWrite := time.Now()
    totalTimeWrite := timeAfterWrite.Sub( timeWrite );
    fmt.Println( "Time write index")
    fmt.Println( totalTimeWrite )

    if newRecs > 0 {
        timeWriteHashes := time.Now()
        writeNewHashes( listDir, newHashes )
        timeAfterWriteHashes := time.Now()
        totalTimeHashes := timeAfterWriteHashes.Sub( timeWriteHashes )
        fmt.Println( "Time write hashes")
        fmt.Println(totalTimeHashes)
    }

    fmt.Printf( "Records: %v New: %v Dupes: %v\n", recs, newRecs, dupes )
    //fmt.Println( index.storage )    

}