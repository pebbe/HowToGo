package main

import (
	"fmt"
	"log"
	"net/http"
)

type handler struct {
	domain string
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Println(r.RemoteAddr, "→", r.Host)
	fmt.Fprintln(w, "Hello from", h.domain)

}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", handler{})
	mux.Handle("localhost/", handler{domain: "localhost"})
	mux.Handle("example.com/", handler{domain: "example.com"})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server listening on port :8080")
	server.ListenAndServe()

}
