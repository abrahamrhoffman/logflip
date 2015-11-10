//// DECLARE PACKAGES ////
package main

//// IMPORT PATHS ////
import (
     "fmt"
     "os"
//     "path/filepath"
//     "flag"
)

//// DECLARE CONSTANTS ////
const dir = "/var/logflip"
const file = "/var/logflip/logflip.conf"
const filetype = ".log"

//// DECLARE FUNCTIONS ////

// Does the constant, 'file' - exist? //
func filechecker() {
     if _, err := os.Stat(file); err == nil {
     } else {
     // Create if 'file' doesn't exist
     fmt.Println("Creating", file)
     os.Mkdir(dir, 0755)
     os.Create(file)
     }
}

//// Find all .log files in the system //
//func findlogs() {
//      d, err := os.Open(pathcheck)
//      if err != nil {
//          fmt.Println(err)
//          os.Exit(1)
//      }
//
//      files, err := d.Readdir(-1)
//      if err != nil {
//          fmt.Println(err)
//          os.Exit(1)
//      }
//
//      fmt.Println("Reading "+ pathcheck)
//
//      for _, file := range files {
//          if file.Mode().IsRegular() {
//              if filepath.Ext(file.Name()) == filetype {
//                fmt.Println(file.Name())
//              }
//          }
//      }
//  }



// func walkpath(path string, f os.FileInfo, err error) error {
//     fmt.Printf("%s with %d bytes\n", path,f.Size())
//     return nil
// }


//// RUN PROGRAM ////
func main() {
    filechecker()
    // Scan '/' recursively for .log files
//    findlogs()


//    flag.Parse()
//    filepath.Walk(pathcheck, walkpath)
    // Append .log files to /var/logflip/logflip.conf

    // Check .log files size. 

    // If over 100MB, check for .log.1

    // If .log.1 exists, erase it
    
    // Create .log.1 as a copy of the original .log
    
    // Erase 
}
