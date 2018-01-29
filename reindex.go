package main 


// import(
//     "ioutil"
//     "strings"
// )

// func reindex( listDir, uploadName string){
//     recs, newRecs, dupes int := 0, 0, 0
//     rUploadName, listDir, _, profile, buffersize = getArgs()
//     rLineChan := make(chan string, buffersize)
//     rHashedLineChan := make(chan string, buffersize)
//     rNewHashChan := make(chan string, buffersize)

//     index = newIndex(listDir + "/GoIndex")
//     index.open()
//     go readUpload(index, uploadName, &lineChan)
//     go forceMd5(&rLineChan, &rHashedLineChan)
//     go checkIndex(&recs, &newRecs, &dupes, index, &rHashedLineChan, &rNewHashChan)
//     for {
//         _, open := <-*newHashChan
//     }
// }

// func readOldFiles(listDir string , &lineChan chan string){
//     index = newIndex(listDir + "/GoIndex")
//     index.open()
//     files := ioutil.Readdir(listDir);
//     for file := range files{
//         if strings.Contains( file.Name ,".md5" ){
//             go readUpload(index, file.Name, &lineChan)
//             go forceMd5(&lineChan, &hashedLineChan)
//             go checkIndex(&recs, &newRecs, &dupes, index, &hashedLineChan, &newHashChan)
//         }
//     }
// }
