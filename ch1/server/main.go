package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var mu sync.Mutex
var count int

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
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
}

func counterHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(size int, cycles int, out io.Writer) {
	const (
		res     = 0.001
		nframes = 64
		delay   = 8
	)

	palette := []color.Color{
		color.Black,
		color.RGBA{0xff, 0x00, 0x00, 0xff}, // red
		color.RGBA{0x00, 0x00, 0xff, 0xff}, // blue
		color.RGBA{0x00, 0x80, 0x00, 0xff}, // green
	}

	rand.Seed(time.Now().UnixNano())
	lineColorIndex := uint8(rand.Intn(len(palette[1:]) + 1))
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), lineColorIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	cycles, err := strconv.Atoi(r.FormValue("cycles"))
	if err != nil {
		cycles = 5
	}
	size, err := strconv.Atoi(r.FormValue("size"))
	if err != nil {
		size = 100
	}
	lissajous(size, cycles, w)
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/count", counterHandler)
	http.HandleFunc("/lissajous", lissajousHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
