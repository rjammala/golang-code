package main

import (
	"archive/zip"
	"log"
	"os"
)

func main() {
	filename, err := os.Create("mytest.zip")

	if err != nil {
		log.Fatal(err)
	}

	w := zip.NewWriter(filename)

	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "this archive contains some text files."},
		{"hello.txt", "hello world? How are you?"},
	}

	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}

		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}
}
