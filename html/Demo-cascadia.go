package main

import (
	"github.com/andybalholm/cascadia"
	"github.com/pebbe/util"
	"golang.org/x/net/html"

	"bytes"
	"fmt"
	"net/http"
)

func main() {
	x := util.CheckErr

	resp, err := http.Get("http://pebbe.tumblr.com")
	x(err)
	doc, err := html.Parse(resp.Body)
	x(err)
	resp.Body.Close()

	sel, err := cascadia.Compile("section.post")
	x(err)

	for i, n := range sel.MatchAll(doc) {
		var b bytes.Buffer
		html.Render(&b, n)
		fmt.Println(i, b.String())
	}

}
