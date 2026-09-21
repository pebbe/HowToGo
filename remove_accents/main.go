package main

import (
	"fmt"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type RT struct {
	rt *unicode.RangeTable
}

func NewRT(rt *unicode.RangeTable) RT {
	return RT{rt: rt}
}

func (t RT) Contains(r rune) bool {
	return unicode.Is(t.rt, r)
}

func main() {
	s := "Yoùr Śtring šđčćžŠĐČĆŽ Ötzi's Nationalität èàì Ĳĳ"
	b := make([]byte, len(s))

	mn := NewRT(unicode.Mn) // Mn: nonspacing marks

	t := transform.Chain(norm.NFKD, runes.Remove(mn), norm.NFKC)
	n, _, e := t.Transform(b, []byte(s), true)
	if e != nil {
		panic(e)
	}

	fmt.Println(string(b[:n]))

}
