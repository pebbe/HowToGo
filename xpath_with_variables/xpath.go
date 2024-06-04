/*

Demo of xpath with variables

*/

package main

import (
	"fmt"
	"unsafe"

	"github.com/jbowtie/gokogiri"
	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/gokogiri/xpath"
	"github.com/pebbe/util"
)

var (
	xmldoc = `<?xml version="1.0"?>
<animals>
  <chimp name="aap">
    <food>banaan</food>
  </chimp>
  <cat name="mies">
    <food>ansjovis</food>
  </cat>
  <chimp name="aapje">
    <food>apenootjes</food>
  </chimp>
</animals>
`
	x = util.CheckErr
)

type varScope struct {
	vars map[string]interface{}
}

func (v *varScope) ResolveVariable(local, namespace string) interface{} {
	return v.vars[local]
}

func (v *varScope) IsFunctionRegistered(a, b string) bool {
	return false
}

func (v *varScope) ResolveFunction(a, b string) xpath.XPathFunction {
	return nil
}

func main() {
	doc, err := gokogiri.ParseXml([]byte(xmldoc))
	x(err)
	root := doc.Root()
	chimps, err := root.Search(`//chimp`)
	x(err)
	cats, err := root.Search(`//cat`)
	x(err)

	foods, err := root.SearchWithVariables(
		`$chimps/food | $cats/food`,
		&varScope{
			vars: map[string]interface{}{
				"chimps": xml.Nodeset(chimps).ToPointers(),
				"cats":   []unsafe.Pointer{cats[0].NodePtr()},
			},
		})
	x(err)
	for _, food := range foods {
		fmt.Println(food)
	}

	doc.Free()
}
