package main

import "fmt"
import "io/ioutil"
import "os"
import "log"
import "path/filepath"
import "strings"

func walkFunction(path string, info os.FileInfo, err error) error {
	var errorVar error
	if info.IsDir() == false {
		bytes, err := ioutil.ReadFile(path)
		s := string(bytes)
		if err != nil {
			errorVar = err
		}
		if strings.Contains(s, "walkFunction") {
			fmt.Println(path)
		}
	
	}
	return errorVar
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("You must supply a directory")
	}
	directory := os.Args[1]
	err := filepath.Walk(directory, walkFunction)
	if err != nil {
		log.Fatal(err)
	}

}
