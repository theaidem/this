package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var port = flag.String("port", "1313", "Your custom port")

func main() {
	flag.Parse()
	if dir, err := os.Getwd(); err == nil {
		log.Println("Current path: ", dir)
		http.Handle("/", http.FileServer(http.Dir(dir)))
		log.Println(fmt.Sprintf("Server listen http://localhost:%s", (*port)))
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", (*port)), nil))
	} else {
		log.Fatal(err)
	}
}
