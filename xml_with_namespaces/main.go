package main

import (
	// "github.com/nbio/xml"
	"encoding/xml"
	"fmt"
)

type DemoT struct {
	//XMLName xml.Name `xml:"http://example.com/ demo"`
	XMLName xml.Name `xml:"demo"`
	Text    string   `xml:"http://example.com/ text"`
	TextAlt string   `xml:"http://example.com/alt text"`
}

var (
	doc = `<?xml version="1.0" encoding="UTF-8"?>
<demo xmlns="http://example.com/" xmlns:alt="http://example.com/alt">
  <text>foo</text>
  <alt:text>bar</alt:text>
</demo>
`
)

func main() {

	fmt.Println(doc)

	var demo DemoT
	err := xml.Unmarshal([]byte(doc), &demo)

	fmt.Printf("%v\n%#v\n\n", err, demo)

	b, err := xml.MarshalIndent(demo, "", "  ")
	fmt.Printf("%v\n%s\n", err, b)
}
