package main

import (
	"golang.org/x/text/collate"
	"golang.org/x/text/language"

	"bytes"
	"fmt"
)

func main() {

	s := []string{"cab", "Cab", "cáb", "dab"}

	for _, lang := range []string{"nl"} {

		ln := language.Make(lang)

		col := collate.New(ln)
		col.SortStrings(s)

		fmt.Println(s)

		var b collate.Buffer
		for _, e := range s {
			b.Reset()
			k := col.KeyFromString(&b, e)
			fmt.Println(e, "\t", hex(k))
		}
	}

}

func hex(b []byte) string {
	var buf bytes.Buffer
	n := len(b)
	skip := false
	for i := 0; i < n; i++ {
		if skip {
			skip = false
			continue
		}
		if i%2 == 0 && i < n-1 && b[i] == 0 && b[i+1] == 0 {
			buf.WriteString("!")
			skip = true
		} else if b[i] == 0 {
			buf.WriteString(".")
		} else {
			buf.WriteString(fmt.Sprintf("%02X", b[i]))
		}
	}
	return buf.String()
}
