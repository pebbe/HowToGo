package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", test)
	http.ListenAndServe(":8099", nil)
}

func test(w http.ResponseWriter, r *http.Request) {

	var ch <-chan bool
	if f, ok := w.(http.CloseNotifier); ok {
		ch = f.CloseNotify()
		fmt.Println("ok")
	} else {
		ch = make(<-chan bool)
		fmt.Println("Geen CloseNotify")
	}

	w.Header().Set("Content-type", "text/plain")
	for i := 0; i < 100; i++ {
		s := strings.Repeat(fmt.Sprintf("%02d", i), 40)
		n, err := fmt.Fprintln(w, s)
		fmt.Println(i, n, err)
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
		time.Sleep(time.Second)
		select {
		case <-ch:
			fmt.Println("Closed")
			return
		default:
		}
	}
}
