package main

import (
	"log"
	"net/http"
	"os"
)

var port = ":1313"

func main() {
	if dir, err := os.Getwd(); err == nil {
		log.Println("Current path: ", dir)
		http.Handle("/", http.FileServer(http.Dir(dir)))
		log.Println("Server listen localhost", port)
		log.Fatal(http.ListenAndServe(port, nil))
	} else {
		log.Fatal(err)
	}
}
