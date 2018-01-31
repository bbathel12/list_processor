package main 

import(
    "archive/zip"
    "bufio"
    _"bytes"
    "fmt"
    "io"
    "os"
)

var zipname string = "list";

func zipper( hashes <-chan string, listDir string ){
    
    flatFileName := writeFlatFile( hashes, listDir )
    err := createZipFile( flatFileName, listDir )
    if err != nil{
        panic( err )
    }
    //remove flat file to save a little space
    os.Remove(flatFileName)
}

func unzipper(){

}

func createZipFile(flatFileName , listDir string) error {
    filename := fmt.Sprintf("%v%c%v.zip", listDir, os.PathSeparator, zipname)
    newfile, err := os.Create(filename)
    newfile.Chmod( os.ModePerm )
    if err != nil {
        return err
    }
    defer newfile.Close()

    zipWriter := zip.NewWriter(newfile)
    defer zipWriter.Close()

    zipfile, err := os.Open(flatFileName)
    if err != nil {
        return err
    }
    defer zipfile.Close()

    // Get the file information
    info, err := zipfile.Stat()
    if err != nil {
        return err
    }

    header, err := zip.FileInfoHeader(info)
    if err != nil {
        return err
    }

    // Change to deflate to gain better compression
    // see http://golang.org/pkg/archive/zip/#pkg-constants
    header.Method = zip.Deflate

    writer, err := zipWriter.CreateHeader(header)
    if err != nil {
        return err
    }
    _, err = io.Copy(writer, zipfile)
    if err != nil {
        return err
    }
    
    return nil
}

func writeFlatFile( hashes <-chan string, listDir string ) (filename string){
    // I know this is stupid but write everything to a real file
    filename = fmt.Sprintf("%v%c%v.txt", listDir, os.PathSeparator, zipname)
    file, _ := os.Create(filename);
    chmodErr := file.Chmod( os.ModePerm )
    if chmodErr != nil{
        panic(chmodErr)
    }
    writer := bufio.NewWriter(file);

    for hashedLine := range hashes{
        //fmt.Printf(".")
        writer.Write( []byte( hashedLine + "\n" ) );
    }

    writer.Flush()
    file.Close()
    return
}