package main

import (
	"math"

	"github.com/adriffaud/ray-tracer-challenge/internal"
)

func main() {
	c := internal.NewCanvas(500, 500)
	p := internal.Point{X: 0, Y: 0, Z: 1}
	translate := internal.Translation(250, 0, 250)
	scale := internal.Scaling(100, 0, 100)

	for h := 0; h < 12; h++ {
		r := internal.RotationY(float64(h) * math.Pi / 6)
		transform := translate.Multiply(scale.Multiply(r))
		p2 := p.MultiplyMatrix(transform)
		c.WritePixel(int(p2.X), c.Height-int(p2.Z), internal.Color{R: 0, G: 1, B: 0})
	}

	internal.Export(c, "./clock.png")
}
