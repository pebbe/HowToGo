/*

Demo of using XPATH with name spaces

*/

package main

import (
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"

	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	data, err := ioutil.ReadFile("xpath.xml")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := xml.Parse(data, nil, nil, 0, xml.DefaultEncodingBytes)
	if err != nil {
		log.Fatal(err)
	}
	defer doc.Free()

	xp := doc.DocXPathCtx()
	xp.RegisterNamespace("folia", "http://ilk.uvt.nl/folia")

	fmt.Println("\nAll sentences with all words:\n")

	xps := xpath.Compile("//folia:s")
	xpw := xpath.Compile("folia:w/folia:t")

	ss, err := doc.Root().Search(xps)
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range ss {
		fmt.Println(s.Attr("id"))
		ww, err := s.Search(xpw)
		if err != nil {
			log.Fatal(err)
		}
		for _, w := range ww {
			fmt.Println("\t" + w.Parent().Attr("id") + "  \t" + w.Content())
		}
	}

	fmt.Println("\nSearch for specific sentence:\n")
	n, err := doc.Root().Search(`//folia:s[@xml:id="WR-P-E-E-0000000020.head.4.s.2"]`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)

	fmt.Println("\nSearch for sentence with specific word:\n")
	n, err = doc.Root().Search(`//folia:w[@xml:id="WR-P-E-E-0000000020.head.4.s.2.w.2"]`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n[0].Parent())

}
