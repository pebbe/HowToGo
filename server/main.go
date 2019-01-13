package main

import (
	"log"
	"net/http"
)

type server struct {
	router *http.ServeMux
}

func main() {
	s := &server{}
	s.routes()
	log.Fatal(http.ListenAndServe(":8666", s.router))
}
