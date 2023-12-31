package internal

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

type Canvas struct {
	Pixels [][]Color
	Width  int
	Height int
}

// WritePixel sets the given color to the x/y coordinates in the canvas.
func (c *Canvas) WritePixel(x, y int, color Color) {
	c.Pixels[y][x] = color
}

// PixelAt returns the pixel color at the given canvas coordinates.
func (c *Canvas) PixelAt(x, y int) Color {
	return c.Pixels[y][x]
}

// NewCanvas initialize a new canvas with every pixel set with a black color.
func NewCanvas(w, h int) *Canvas {
	black := Color{R: 0, G: 0, B: 0}

	p := make([][]Color, h)
	for x := 0; x < h; x++ {
		p[x] = make([]Color, w)
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			p[y][x] = black
		}
	}

	return &Canvas{Width: w, Height: h, Pixels: p}
}

// Export exports the given canvas as a PNG image a the given path.
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

func scaleColor(c Color) color.Color {
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
