package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

type ThingListT []*ThingT

type DocT struct {
	XMLName xml.Name   `xml:"doc"`
	Thing   ThingListT `xml:"thing,omitempty"`
}

type ThingT struct {
	ID     string     `xml:"id,attr,omitempty"`
	Foo    string     `xml:"foo,attr,omitempty"`
	Bar    string     `xml:"bar,attr,omitempty"`
	Qin    int        `xml:"qin,attr,omitempty"`
	Thing  ThingListT `xml:"thing,omitempty"`
	parent *ThingT
}

var (
	data = `<?xml version="1.0"?>
<doc>
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
	for _, thing := range doc.Thing {
		setParents(thing)
	}

	fmt.Println(data)

	////////////////////////////////

	fmt.Println(`xpath: /doc/thing[@foo="yes" and not(@bar)]`)

	if doc.Thing.Test(Foo("yes"), Bar("")) {
		fmt.Println("found")
	} else {
		fmt.Println("not found")
	}
	fmt.Println()

	fmt.Println(`xpath: /doc/thing[@foo="yes" and @bar="yes"]`)
	if things := doc.Thing.List(Foo("yes"), Bar("yes")); Any(things) {
		for _, d := range things {
			fmt.Println("found:")
			fmt.Println(d)
		}
	} else {
		fmt.Println("not found")
	}
	fmt.Println("as simple Go")
	for _, thing := range doc.Thing {
		if thing.Foo == "yes" && thing.Bar == "yes" {
			fmt.Println("found:")
			fmt.Println(thing)
		}
	}
	fmt.Println()

	////////////////////////////////

	fmt.Println(`xpath: /doc/thing[@bar]/thing[not(@bar)]`)

	if things := doc.Thing.List(NotBar("")).Thing().List(Bar("")); Any(things) {
		for _, d := range things {
			fmt.Println("found:")
			fmt.Println(d)
		}
	} else {
		fmt.Println("not found")
	}
	fmt.Println("as simple Go")
	for _, thing := range doc.Thing {
		if thing.Bar != "" {
			for _, thing2 := range thing.Thing {
				if thing2.Bar == "" {
					fmt.Println("found:")
					fmt.Println(thing2)
				}
			}
		}
	}
	fmt.Println()

	////////////////////////////////

	fmt.Println(`xpath: /doc/thing[@bar="yes" or @qin > 1]`)

	if things := doc.Thing.List(Or(Bar("yes"), QinGreater(1))); Any(things) {
		for _, d := range things {
			fmt.Println("found:")
			fmt.Println(d)
		}
	} else {
		fmt.Println("not found")
	}
	fmt.Println()

	fmt.Println(`xpath: /doc//thing[@qin && (@qin < 10 or @qin > 100)]`)
	if things := doc.Thing.DescendantOrSelfThing().List(NotQin(0), Or(QinLess(10), QinGreater(100))); Any(things) {
		for _, d := range things {
			fmt.Println("found:")
			fmt.Println(d)
		}
	} else {
		fmt.Println("not found")
	}
	fmt.Println()

	////////////////////////////////

	fmt.Println(`xpath: /doc//thing[@qin = 333]/..`)

	if things := doc.Thing.DescendantOrSelfThing().List(Qin(333)).Parent(); Any(things) {
		for _, d := range things {
			fmt.Println("found:")
			fmt.Println(d)
		}
	} else {
		fmt.Println("not found")
	}
	fmt.Println("as simple Go")
	var f1 func(thing ThingListT)
	f1 = func(thing ThingListT) {
		for _, th := range thing {
			if th.Qin == 333 && th.parent != nil {
				fmt.Println("found:")
				fmt.Println(th.parent)
			}
			f1(th.Thing)
		}
	}
	f1(doc.Thing)
	fmt.Println()

	////////////////////////////////

	fmt.Println(`xpath: /doc//thing[@qin = 1 and thing[@qin < 0]]`)

	if things := doc.Thing.DescendantOrSelfThing().List(Qin(1), Thing(QinLess(0))); Any(things) {
		for _, d := range things {
			fmt.Println("found:")
			fmt.Println(d)
		}
	} else {
		fmt.Println("not found")
	}
	fmt.Println()

}

func setParents(thing *ThingT) {
	for _, t := range thing.Thing {
		t.parent = thing
		setParents(t)
	}
}

func (t *ThingT) String() string {
	b, err := xml.MarshalIndent(t, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(
		strings.Replace(string(b), "ThingT", "thing", -1),
		"></thing>", " />", -1)
}

func Any(things ThingListT) bool {
	return len(things) > 0
}

// methods on ThingListT to get ThingListT or test

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

func (things ThingListT) ThingTest(t ...func(*ThingT) bool) bool {
LOOP:
	for _, thing := range things {
		for _, thing2 := range thing.Thing {
			for _, test := range t {
				if !test(thing2) {
					continue LOOP
				}
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
	// TODO: remove doubles
	return tt
}

func (things ThingListT) DescendantOrSelfThing() ThingListT {
	desc := ThingListT(make([]*ThingT, 0))
	descendantOrSelfThingHelper(things, &desc)
	// TODO: remove doubles
	return desc
}

func descendantOrSelfThingHelper(things ThingListT, desc *ThingListT) {
	for _, thing := range things {
		*desc = append(*desc, thing)
		descendantOrSelfThingHelper(thing.Thing, desc)
	}
}

func (things ThingListT) Thing() ThingListT {
	tt := make([]*ThingT, 0)
	for _, thing := range things {
		tt = append(tt, thing.Thing...)
	}
	// TODO: remove doubles
	return tt
}

func (things ThingListT) HasParent() bool {
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
	// TODO: remove doubles
	return parents
}

// create test closures

func Or(t ...func(*ThingT) bool) func(t *ThingT) bool {
	return func(tt *ThingT) bool {
		for _, test := range t {
			if test(tt) {
				return true
			}
		}
		return false
	}
}

func Foo(v string) func(*ThingT) bool {
	return func(t *ThingT) bool {
		return t.Foo == v
	}
}

func Bar(v string) func(*ThingT) bool {
	return func(t *ThingT) bool {
		return t.Bar == v
	}
}

func NotBar(v string) func(*ThingT) bool {
	return func(t *ThingT) bool {
		return t.Bar != v
	}
}

func Qin(v int) func(*ThingT) bool {
	return func(t *ThingT) bool {
		return t.Qin == v
	}
}

func NotQin(v int) func(*ThingT) bool {
	return func(t *ThingT) bool {
		return t.Qin != v
	}
}

func QinLess(v int) func(*ThingT) bool {
	return func(t *ThingT) bool {
		return t.Qin < v
	}
}

func QinGreater(v int) func(*ThingT) bool {
	return func(t *ThingT) bool {
		return t.Qin > v
	}
}

func Thing(t ...func(*ThingT) bool) func(t *ThingT) bool {
	return func(thing *ThingT) bool {
		for _, test := range t {
			for _, thing2 := range thing.Thing {
				if test(thing2) {
					return true
				}
			}
		}
		return false
	}
}
