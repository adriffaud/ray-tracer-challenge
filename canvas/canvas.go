package canvas

import (
	"github.com/adriffaud/ray-tracer-challenge/color"
)

type Canvas struct {
	Pixels [][]*color.Color
	Width  int
	Height int
}

func NewCanvas(w, h int) *Canvas {
	black := color.NewColor(0, 0, 0)

	p := make([][]*color.Color, h)
	for x := 0; x < h; x++ {
		p[x] = make([]*color.Color, w)
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			p[y][x] = black
		}
	}

	return &Canvas{Width: w, Height: h, Pixels: p}
}

func (c *Canvas) WritePixel(x, y int, color *color.Color) {
	c.Pixels[y][x] = color
}

func (c *Canvas) PixelAt(x, y int) *color.Color {
	return c.Pixels[y][x]
}
