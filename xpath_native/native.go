package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type ThingListT []*ThingT

type DocT struct {
	XMLName xml.Name   `xml:"doc"`
	Things  ThingListT `xml:"thing,omitempty"`
}

type ThingT struct {
	ID     string     `xml:"id,attr,omitempty"`
	Foo    string     `xml:"foo,attr,omitempty"`
	Bar    string     `xml:"bar,attr,omitempty"`
	Qin    int        `xml:"qin,attr,omitempty"`
	Things ThingListT `xml:"thing,omitempty"`
	parent *ThingT
}

var (
	data = `<doc>
  <thing id="A" foo="yes" bar="yes">
    <thing id="B" foo="yes" />
  </thing>
  <thing id="C" qin="1">
    <thing id="D" qin="22">
      <thing id="E" qin="333" />
    </thing>
    <thing id="F" qin="4444" />
  </thing>
  <thing id="G" qin="1">
    <thing id="H" qin="-10" />
  </thing>
</doc>
`
)

func main() {

	var doc DocT
	err := xml.Unmarshal([]byte(data), &doc)
	if err != nil {
		log.Fatal(err)
	}
	for _, thing := range doc.Things {
		setParents(thing)
	}

	// $doc/thing[@foo="yes" and not(@bar)]
	if doc.Things.Test(foo("yes"), bar("")) {
		fmt.Println("found")
		for _, d := range doc.Things.List(foo("yes"), bar("")) {
			fmt.Println(d)
		}
	} else {
		fmt.Println("not found")
	}

	// $doc/thing[@foo="yes" and @bar="yes"]
	if doc.Things.Test(foo("yes"), bar("yes")) {
		fmt.Println("found")
		for _, d := range doc.Things.List(foo("yes"), bar("yes")) {
			fmt.Println(d)
		}
	} else {
		fmt.Println("not found")
	}

	// $doc/thing[@bar="yes" or @qin > 1]
	if doc.Things.Test(or(bar("yes"), qinGreater(1))) {
		fmt.Println("found")
		for _, d := range doc.Things.List(or(bar("yes"), qinGreater(1))) {
			fmt.Println(d)
		}
	} else {
		fmt.Println("not found")
	}

	// $doc//thing[@qin && (@qin < 10 or @qin > 100)]
	if doc.Things.DescendantOrSelfThing().Test(notQin(0), or(qinLess(10), qinGreater(100))) {
		fmt.Println("found")
		for _, d := range doc.Things.DescendantOrSelfThing().List(notQin(0), or(qinLess(10), qinGreater(100))) {
			fmt.Println(d)
		}
	} else {
		fmt.Println("not found")
	}

	// $doc//thing[@qin = 333]/..
	if doc.Things.DescendantOrSelfThing().List(qin(333)).hasParent() {
		fmt.Println("found")
		for _, d := range doc.Things.DescendantOrSelfThing().List(qin(333)).Parent() {
			fmt.Println(d)
		}
	} else {
		fmt.Println("not found")
	}

	// $doc//thing[@qin = 1 and thing[@qin < 0]]
	if doc.Things.DescendantOrSelfThing().Test(qin(1), Thing(qinLess(0))) {
		fmt.Println("found")
		for _, d := range doc.Things.DescendantOrSelfThing().List(qin(1), Thing(qinLess(0))) {
			fmt.Println(d)
		}
	} else {
		fmt.Println("not found")
	}

}

func setParents(thing *ThingT) {
	for _, t := range thing.Things {
		t.parent = thing
		setParents(t)
	}
}

func (t *ThingT) String() string {
	b, err := xml.MarshalIndent(t, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func (things ThingListT) Test(t ...func(*ThingT) bool) bool {
LOOP:
	for _, thing := range things {
		for _, test := range t {
			if !test(thing) {
				continue LOOP
			}
		}
		return true
	}
	return false
}

func (things ThingListT) List(t ...func(*ThingT) bool) ThingListT {
	tt := make([]*ThingT, 0)
LOOP:
	for _, thing := range things {
		for _, test := range t {
			if !test(thing) {
				continue LOOP
			}
		}
		tt = append(tt, thing)
	}
	return tt
}

func (things ThingListT) DescendantOrSelfThing() ThingListT {
	desc := ThingListT(make([]*ThingT, 0))
	DescendantOrSelfThingHelper(things, &desc)
	return desc
}

func DescendantOrSelfThingHelper(things ThingListT, desc *ThingListT) {
	for _, thing := range things {
		*desc = append(*desc, thing)
		DescendantOrSelfThingHelper(thing.Things, desc)
	}
}

func (things ThingListT) hasParent() bool {
	for _, thing := range things {
		if thing.parent != nil {
			return true
		}
	}
	return false
}

func (things ThingListT) Parent() ThingListT {
	parents := ThingListT(make([]*ThingT, 0))
	for _, thing := range things {
		if thing.parent != nil {
			parents = append(parents, thing.parent)
		}
	}
	return parents
}

func or(t ...func(*ThingT) bool) func(t *ThingT) bool {
	return func(tt *ThingT) bool {
		for _, test := range t {
			if test(tt) {
				return true
			}
		}
		return false
	}
}

func foo(v string) func(*ThingT) bool {
	return func(t *ThingT) bool {
		return t.Foo == v
	}
}

func bar(v string) func(*ThingT) bool {
	return func(t *ThingT) bool {
		return t.Bar == v
	}
}

func qin(v int) func(*ThingT) bool {
	return func(t *ThingT) bool {
		return t.Qin == v
	}
}

func notQin(v int) func(*ThingT) bool {
	return func(t *ThingT) bool {
		return t.Qin != v
	}
}

func qinLess(v int) func(*ThingT) bool {
	return func(t *ThingT) bool {
		return t.Qin < v
	}
}

func qinGreater(v int) func(*ThingT) bool {
	return func(t *ThingT) bool {
		return t.Qin > v
	}
}

func Thing(t ...func(*ThingT) bool) func(t *ThingT) bool {
	return func(thing *ThingT) bool {
		for _, test := range t {
			for _, thing2 := range thing.Things {
				if test(thing2) {
					return true
				}
			}
		}
		return false
	}
}
