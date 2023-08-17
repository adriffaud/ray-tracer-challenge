package canvas

import (
	"reflect"
	"testing"

	"github.com/adriffaud/ray-tracer-challenge/color"
)

func TestCanvasCreation(t *testing.T) {
	c := NewCanvas(10, 20)

	if c.Width != 10 || c.Height != 20 {
		t.Fatalf("expected a 10x20 canvas. got=%+v", c)
	}

	black := color.Color{R: 0, G: 0, B: 0}

	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			p := c.Pixels[y][x]

			if !reflect.DeepEqual(p, black) {
				t.Fatalf("expected pixel to be black. got=%+v", p)
			}
		}
	}
}

func TestPixelWrite(t *testing.T) {
	c := NewCanvas(10, 20)
	red := color.Color{R: 1, G: 0, B: 0}

	c.WritePixel(2, 3, red)
	actual := c.PixelAt(2, 3)
	if !reflect.DeepEqual(actual, red) {
		t.Fatalf("expected pixel to be red. got=%+v", actual)
	}
}
