/*

   Voor simpele dingen, zie Demo-cascadia.go

*/

package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/pebbe/util"

	"fmt"
	"net/http"
)

func main() {
	x := util.CheckErr

	resp, err := http.Get("http://pebbe.tumblr.com")
	x(err)
	doc, err := goquery.NewDocumentFromResponse(resp)
	x(err)
	resp.Body.Close()
	doc.Find("div.post").EachWithBreak(func(i int, s *goquery.Selection) bool {
		h, err := s.Html()
		x(err)
		fmt.Println(">>>", h)
		return true
	})

}
