package main

import (
    "os"
    "log"
    "errors"
    "path/filepath"
    "mime"
    "sync"
)

var wg sync.WaitGroup

func checkErr(err error){
    if err != nil {
        log.Println(err)
    }
}

func Sort(path string) error {
    
    
    handle, err := os.Lstat(path)
    
    if !handle.IsDir() {
        log.Print("It's not a directory")
        return errors.New("error")
    }
    
    directories := []string{"Images", "Music", "Videos", "Documents"}
    
    for _, subfolder := range directories {
        os.Mkdir(filepath.Join(path, subfolder), 0777)
    }
    
    dir, err := os.Open(path)
    fi, err := dir.Readdir(100)
    for _, file := range fi {
        wg.Add(1)
        go mapToDir(path, file.Name())
    }
        
    wg.Wait()
    return err
    
}

func mapToDir(base, name string) error {
    defer wg.Done()
    
    ext := filepath.Ext(name)
    
    log.Println(mime.TypeByExtension(ext))
    
    return nil
}

//Creates if directory doesn't already exist
//func CreateINE(dir string){
//    //Should lock other goroutines
//    _, err := os.Stat(dir)
//    
//    
//    if os.IsNotExist(err) {
//        os.Mkdir(dir, 0777)
//    }
//    
//}
                
    
    
    
    
    
    
    