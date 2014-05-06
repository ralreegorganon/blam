package main

import (
	"flag"
	"log"
	"net/http"
)

var port = flag.String("port", "8080", "TCP port to bind to")
var directory = flag.String("directory", ".", "Directory to serve")

func main() {
	flag.Parse()
	err := http.ListenAndServe(":"+*port, http.FileServer(http.Dir(*directory)))
	if err != nil {
		log.Fatal(err)
	}
}
