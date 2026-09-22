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
	s := "Yoùr Śtring šđčćžŠĐČĆŽ Ötzi's Nationalität èàì æÆœŒ Ĳĳ ﬁﬂﬃ"
	fmt.Println("ori: ", s)

	b := make([]byte, len(s))

	mn := NewRT(unicode.Mn) // Mn: nonspacing marks

	t1 := transform.Chain(norm.NFD, runes.Remove(mn), norm.NFC)
	t2 := transform.Chain(norm.NFKD, runes.Remove(mn), norm.NFKC)

	n, _, e := t1.Transform(b, []byte(s), true)
	if e != nil {
		panic(e)
	}
	fmt.Println("NFD: ", string(b[:n]))

	n, _, e = t2.Transform(b, []byte(s), true)
	if e != nil {
		panic(e)
	}
	fmt.Println("NFKD:", string(b[:n]))
}
