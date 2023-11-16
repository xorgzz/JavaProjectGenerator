package main
import (
    "fmt"
    "os"
    "strings"
)
// Function that checks if directory of project exists
// e.g. project dir=/home/user/project1 function checks if /home/user exists
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
    var pack string

    if err == nil {
        fmt.Printf("Project directory: ")
        fmt.Scan(&finalLoc)
        if finalLoc[0] != '/' {
           finalLoc = projLoc + "/" + finalLoc 
        }
    
        fmt.Print("Project Alias: ")
        fmt.Scan(&pack)
        
        tmp := strings.SplitAfter(finalLoc, "/")
        path := strings.Join(tmp[0:len(tmp)-1], "")
        if exists(path) {
            fmt.Println("[!] Creating Project")
            os.Mkdir(finalLoc, 0755)
            srcPath := finalLoc + "/src"
            os.Mkdir(srcPath, 0755)
            os.Mkdir(srcPath+"/"+pack, 0755)
            outPath := finalLoc + "/out"
            os.Mkdir(outPath, 0755)
            os.WriteFile(srcPath+"/"+pack+"/Main.java", []byte("package "+pack+";\n\npublic class Main {\n\tpublic static void main(String[] args) {\n\n\t}\n}"), 0666)
            os.WriteFile(finalLoc+"/build.sh", []byte("#!/usr/bin/sh\n\njavac -d "+outPath+" "+srcPath+"/"+pack+"/*.java && cd out && java "+pack+".Main"), 0766)
        } else {
            fmt.Println("[!] Directory doesn't exist")
        }
    } else {
        fmt.Println(err)
    }
    
}
