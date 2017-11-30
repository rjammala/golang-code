package main

import (
	"io/ioutil"
	"log"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
)

func main() {
	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		file := r.FormValue("file")
		b, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		j, err := simplejson.NewJson(b)
		if err != nil {
			log.Fatal(err)
		}
		result, err := j.EncodePretty()
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
		w.Write([]byte("\n"))
	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
