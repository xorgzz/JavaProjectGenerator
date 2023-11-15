package main
import (
    "fmt"
    "os"
    "strings"
)

func exists(path string) (bool) {
    info, err := os.Stat(path)
    if !os.IsNotExist(err) && info.IsDir() {
        return true
    }
    return false
}

func main () {
    fmt.Println("Java Project Creator")
    projLoc, err := os.Getwd()
    var finalLoc string
    if err == nil {
        fmt.Printf("Current project directory %s\n", projLoc)
        fmt.Printf("Project directory: ")
        fmt.Scan(&finalLoc)
        if finalLoc[0] != '/' {
           finalLoc = projLoc + "/" + finalLoc 
        }

        tmp := strings.SplitAfter(finalLoc, "/")
        path := strings.Join(tmp[0:len(tmp)-1], "")
        if exists(path) {
            fmt.Println("[!] Creating Project")
            os.Mkdir(finalLoc, 0755)
            srcPath := finalLoc + "/src"
            os.Mkdir(srcPath, 0755)
            outPath := finalLoc + "/out"
            os.Mkdir(outPath, 0755)
            os.WriteFile(srcPath+"/Main.java", []byte("import java.util.*\n\nclass Main {\n\tpublic static void main(String[], args) {\n\n\t}\n}"), 0666)
            os.WriteFile(finalLoc+"/build.sh", []byte("#!/usr/bin/sh\n\njavac -d "+outPath+" "+srcPath+"/*.java && cd out && java out.Main"), 0766)
        } else {
            fmt.Println("[!] Directory doesn't exist")
        }
    } else {
        fmt.Println(err)
    }
    
}
