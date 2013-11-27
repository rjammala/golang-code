package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var patternToGrep string

func printMatch(path string, s string, patternToGrep string, sc chan string) {
	if strings.Contains(s, patternToGrep) {
		sc <- path
	} else {
		sc <- ""
	}
}

func walkFunction(path string, info os.FileInfo, err error) error {
	var errorVar error

	sc := make(chan string)

	if info.IsDir() == false {
		bytes, err := ioutil.ReadFile(path)
		s := string(bytes)
		if err != nil {
			errorVar = err
		}

		go printMatch(path, s, patternToGrep, sc)
		result := <-sc
		if result != "" {
			fmt.Println(result)
		}
	}
	return errorVar
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("You must supply a directory and a string to grep")
	}
	directory := os.Args[1]
	patternToGrep = os.Args[2]
	err := filepath.Walk(directory, walkFunction)
	if err != nil {
		log.Fatal(err)
	}
}
