package main

import (
	"math"

	"github.com/adriffaud/ray-tracer-challenge/canvas"
	"github.com/adriffaud/ray-tracer-challenge/color"
	"github.com/adriffaud/ray-tracer-challenge/matrix"
	"github.com/adriffaud/ray-tracer-challenge/tuple"
)

func main() {
	c := canvas.NewCanvas(500, 500)
	p := tuple.Point{XVal: 0, YVal: 0, ZVal: 1}
	translate := matrix.Translation(250, 0, 250)
	scale := matrix.Scaling(100, 0, 100)

	for h := 0; h < 12; h++ {
		r := matrix.RotationY(float64(h) * math.Pi / 6)
		transform := translate.Multiply(scale.Multiply(r))
		p2, err := transform.MultiplyTuple(&p)
		if err != nil {
			panic(err)
		}
		c.WritePixel(int(p2.X()), c.Height-int(p2.Z()), color.Color{R: 0, G: 1, B: 0})
	}

	canvas.Export(c, "./clock.png")
}
