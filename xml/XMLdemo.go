/*

How to handle unknown attributes

*/

package main

import (
	"github.com/kr/pretty"

	"encoding/xml"
	"fmt"
)

type Xml struct {
	XMLName xml.Name `xml:"demo"`
	Item    ItemType `xml:"item"`
}

type ItemType struct {
	Aap  string `xml:"aap,attr"`
	Noot string `xml:"noot,attr"`
	Mies string `xml:"mies,attr"`
	// Other map[string]string `xml:"-"`
	other map[string]string
	Items []ItemType `xml:"item"`
}

// ter voorkoming van oneindige recursie
type ItemTT ItemType

var (
	ItemTags = []string{
		"aap",
		"noot",
		"mies",
	}
	keyTags = make(map[string]bool)
)

func init() {
	for _, tag := range ItemTags {
		keyTags[tag] = true
	}
}

func (x *ItemType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if n := attr.Name.Local; !keyTags[n] {
			if x.other == nil {
				x.other = make(map[string]string)
			}
			x.other[n] = attr.Value
		}
	}
	return d.DecodeElement((*ItemTT)(x), &start)
}

func main() {
	data := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<demo>
  <item aap="baviaan" noot="walnoot" mies="poes" wim="goudvis">
    <item aap="chimp" noot="pinda" mies="mispoes" wim="guppy" />
  </item>
</demo>
`)
	var doc Xml
	fmt.Println(xml.Unmarshal(data, &doc))
	pretty.Println(doc)
}
