package main

import (
	"image"
	"image/color"
	"image/gif"
	"os"
)

var palette = []color.Color{color.Black, color.White}
var nframes = 64

func main() {
	anim := gif.GIF{LoopCount: nframes}
	rect := image.Rect(0, 0, 1000, 1000)
	img := image.NewPaletted(rect, palette)

	for i := nframes; i > 0; i-- {
		img.SetColorIndex(i, i, 1)
		anim.Delay = append(anim.Delay, 8)
		anim.Image = append(anim.Image, img)
	}
	f, err := os.Create("1.gif")
	if err != nil {

	}
	gif.EncodeAll(f, &anim)
	f.Close()
}
