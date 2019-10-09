package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}, color.RGBA{0xFF, 0x00, 0x00, 0xFF}}

const (
	blackIndex = 0
	greenIndex = 1
	redIndex   = 2
)

type lParams struct {
	cycles int
	size   int
	delay  int
}

func lissajous(out io.Writer, params *lParams) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	var cycles = 5 // Default value
	if params.cycles != 0 {
		cycles = params.cycles
	}

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			var index uint8
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			if x > 0.3 {
				index = redIndex
			} else {
				index = greenIndex
			}

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), index)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
