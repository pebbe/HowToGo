package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
)

type context struct {
	w  http.ResponseWriter
	r  *http.Request
	db string
}

func (s *server) routes() {
	s.router = http.NewServeMux()

	s.handleFunc("/hello", handleHello)
	s.handleFunc("/admin", handleAdmin, isAdmin)
	s.handleFunc("/dbase", handleDB, openDB)
}

func (s *server) handleFunc(url string, handler func(*context), options ...func(*context) (ok bool, defered func())) {
	url = path.Join("/", url)
	s.router.HandleFunc(
		url,
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[%s] %s %s %s", r.Header.Get("X-Forwarded-For"), r.RemoteAddr, r.Method, r.URL)

			if r.URL.Path != url {
				http.NotFound(w, r)
				return
			}

			q := &context{
				w: w,
				r: r,
			}

			for _, option := range options {
				ok, defered := option(q)
				if defered != nil {
					defer defered()
				}
				if !ok {
					return
				}
			}

			log.Print("Running handler for ", url)
			handler(q)
		})
}

// HANDLERS

func handleHello(q *context) {
	fmt.Fprintln(q.w, "Hello")
}

func handleAdmin(q *context) {
	fmt.Fprintln(q.w, "Hello admin")
}

func handleDB(q *context) {
	fmt.Fprintf(q.w, "Database: %v\n", q.db)
}

// HANDLER OPTIONS

func isAdmin(q *context) (ok bool, defered func()) {
	admin := true

	if !admin {
		http.NotFound(q.w, q.r)
		return false, nil
	}

	return true, nil
}

func openDB(q *context) (ok bool, defered func()) {
	q.db = "OPENED DATABASE"
	log.Printf("Opening database: %v", q.db)
	return true, func() {
		log.Printf("Closing database: %v", q.db)
		q.db = "CLOSED DATABASE"
	}
}
