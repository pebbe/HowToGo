package main

import (
	"log"
	"net/http"
)

func main() {
	s := &server{}
	s.routes()
	log.Fatal(http.ListenAndServe(":8666", s.router))
}
