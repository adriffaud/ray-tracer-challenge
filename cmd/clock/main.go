package main

import (
	"math"

	"github.com/adriffaud/ray-tracer-challenge/pkg/canvas"
	"github.com/adriffaud/ray-tracer-challenge/pkg/color"
	"github.com/adriffaud/ray-tracer-challenge/pkg/primitives"
)

func main() {
	c := canvas.NewCanvas(500, 500)
	p := primitives.Point{X: 0, Y: 0, Z: 1}
	translate := primitives.Translation(250, 0, 250)
	scale := primitives.Scaling(100, 0, 100)

	for h := 0; h < 12; h++ {
		r := primitives.RotationY(float64(h) * math.Pi / 6)
		transform := translate.Multiply(scale.Multiply(r))
		p2 := p.MultiplyMatrix(transform)
		c.WritePixel(int(p2.X), c.Height-int(p2.Z), color.Color{R: 0, G: 1, B: 0})
	}

	canvas.Export(c, "./clock.png")
}
