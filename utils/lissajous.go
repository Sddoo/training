package utils

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"time"
	"net/url"
	"net/http"
	"strconv"
)

// var palette = []color.Color{color.Black, color.RGBA{0, 255, 0, 255}}

// const (
// 	blackIndex = 0 // next color in palette
// 	greenIndex = 1
// )

func Lissajous(out io.Writer, r *http.Request) {
	parsedParams, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return
	}
	cycles, err := strconv.ParseFloat(parsedParams.Get("cycles"), 10)
	if err != nil {
		cycles = 5
	}
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	
	var palette = []color.Color{color.Black}
	for i := 0; i < 254; i++ {
		var newColor = []color.Color{color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(255)}}
		palette = append(palette, newColor[0])
	}
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(254)) + 1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
