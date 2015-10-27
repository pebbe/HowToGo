/*

Maak een PNG met precies de grootte van een tekst (plus margin)

De achtergrond is transparant, totdat je het opvult met wit (of een andere kleur)

*/

package main

import (
	"github.com/ungerik/go-cairo"
)

func main() {

	text := "Mói, Groningen! €"
	font := "serif"
	fonti := cairo.FONT_SLANT_ITALIC
	fontw := cairo.FONT_WEIGHT_BOLD
	fontsize := 32.0
	margin := 1.0

	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, 1, 1)
	surface.SelectFontFace(font, fonti, fontw)
	surface.SetFontSize(fontsize)
	te := surface.TextExtents(text)
	x := te.Xbearing
	y := te.Ybearing
	w := te.Width
	h := te.Height
	// dx := te.Xadvance
	// dy := te.Yadvance
	surface.Finish()

	surface = cairo.NewSurface(cairo.FORMAT_ARGB32, int(2*margin+w+.5), int(2*margin+h+.5))
	surface.SelectFontFace(font, fonti, fontw)
	surface.SetFontSize(fontsize)
	surface.MoveTo(margin-x, margin-y)
	surface.ShowText(text)
	surface.WriteToPNG("tekst.png")
	surface.Finish()

}
