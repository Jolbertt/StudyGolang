package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int
var pallete = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}, color.Opaque, color.RGBA{0xFF, 0xAB, 0x00, 0xFF}, color.RGBA{0xFF, 0x00, 0x00, 0xFF}, color.RGBA{0x00, 0x00, 0xFF, 0xFF}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query, err := strconv.ParseFloat(r.URL.Query().Get("cycles"), 64)
		if err != nil {
			log.Print(err)
		}

		if query != 0.0 {
			lissajous(w, query)
		} else {
			lissajous(w, 5.0)
		}
	})
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8088", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(out http.ResponseWriter, cycles float64) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t + 100)
			y := math.Cos(t*freq + phase)
			randColor := rand.Intn(5)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(randColor))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
