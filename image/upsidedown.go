package main

import (
	"image"
	"image/draw"
	_ "image/png"

	"log"
	"os"
)

func main() {
	fp, err := os.Open("image.png")
	x(err)
	img, _, err := image.Decode(fp)
	fp.Close()
	x(err)

	rgba := image.NewRGBA(img.Bounds())
	/*
		if rgba.Stride != rgba.Rect.Size().X*4 {
			x(errors.New("unsupported stride"))
		}
	*/

	// flip image upside down
	ib := img.Bounds()
	width := ib.Max.X
	height := ib.Max.Y
	for h := 0; h < height; h++ {
		r := image.Rect(0, height-1-h, width, height-h)
		draw.Draw(rgba, r, img, image.Point{0, h}, draw.Src)
	}
}

func x(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
