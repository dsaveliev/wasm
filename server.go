package main

import (
	"log"
	"net/http"
	"regexp"
)

var fs = http.FileServer(http.Dir("."))
var wasm = regexp.MustCompile("\\.wasm$")

func fileserver(w http.ResponseWriter, r *http.Request) {
	if wasm.MatchString(r.RequestURI) {
		w.Header().Set("Content-Type", "application/wasm")
	}
	fs.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", fileserver)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
