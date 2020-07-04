package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// var palette = []color.Color{color.White, color.Black}
// var palette = []color.Color{color.White, color.RGBA{0x00, 0xff, 0x00, 0xff}}
var palette = []color.Color{
	color.White,
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
}

const (
	whiteIndex = 0
	// greenIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
		fill    = 0.9
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	no0fColors := len(palette) - 1

	for i := 0; i < nframes; i++ {
		// rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		// img := image.NewPaletted(rect, palette)

		// rect := image.Rect(0, 0, 2*size, 2*size)
		// img := image.NewPaletted(rect, palette)

		colorIndex := uint8(i%no0fColors + 1)
		rect := image.Rect(0, 0, 2*size, 2*size)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
			// img.SetColorIndex(size+int(fill*x*size), size+int(fill*y*size), greenIndex)
			img.SetColorIndex(size+int(fill*x*size), size+int(fill*y*size), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
