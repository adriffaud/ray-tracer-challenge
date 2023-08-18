package canvas

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"

	rtc_color "github.com/adriffaud/ray-tracer-challenge/color"
)

type Canvas struct {
	Pixels [][]rtc_color.Color
	Width  int
	Height int
}

func (c *Canvas) WritePixel(x, y int, color rtc_color.Color) {
	c.Pixels[y][x] = color
}

func (c *Canvas) PixelAt(x, y int) rtc_color.Color {
	return c.Pixels[y][x]
}

func NewCanvas(w, h int) *Canvas {
	black := rtc_color.Color{R: 0, G: 0, B: 0}

	p := make([][]rtc_color.Color, h)
	for x := 0; x < h; x++ {
		p[x] = make([]rtc_color.Color, w)
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			p[y][x] = black
		}
	}

	return &Canvas{Width: w, Height: h, Pixels: p}
}

func Export(c *Canvas, path string) {
	img := image.NewNRGBA(image.Rect(0, 0, c.Width, c.Height))

	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			p := c.Pixels[y][x]
			img.Set(x, y, scaleColor(p))
		}
	}

	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func scaleColor(c rtc_color.Color) color.Color {
	return color.RGBA{
		R: scaleComponent(c.R),
		G: scaleComponent(c.G),
		B: scaleComponent(c.B),
		A: 255,
	}
}

func scaleComponent(c float64) uint8 {
	if c < 0 {
		return 0
	} else if c > 1 {
		return 255
	} else {
		return uint8(math.Ceil(c * 255))
	}
}
