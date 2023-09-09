package main

import (
	"github.com/adriffaud/ray-tracer-challenge/pkg/canvas"
	"github.com/adriffaud/ray-tracer-challenge/pkg/color"
	"github.com/adriffaud/ray-tracer-challenge/pkg/primitives"
	"github.com/adriffaud/ray-tracer-challenge/pkg/shapes"
)

const (
	CANVAS_PIXELS = 100.0
	WALL_Z        = 10
	WALL_SIZE     = 7.0
	PIXEL_SIZE    = WALL_SIZE / CANVAS_PIXELS
	HALF          = WALL_SIZE / 2
)

func main() {
	rayOrigin := primitives.Point{Z: -5}

	c := canvas.NewCanvas(CANVAS_PIXELS, CANVAS_PIXELS)
	color := color.Color{R: 1}
	shape := shapes.Sphere()

	for y := 0; y < CANVAS_PIXELS; y++ {
		worldY := HALF - PIXEL_SIZE*float64(y)
		for x := 0; x < CANVAS_PIXELS; x++ {
			worldX := -HALF + PIXEL_SIZE*float64(x)
			position := primitives.Point{
				X: worldX,
				Y: worldY,
				Z: WALL_Z,
			}
			direction := position.SubPoint(rayOrigin).Normalize()

			r := primitives.Ray{Origin: rayOrigin, Direction: direction}
			xs := shape.Intersect(r)

			if xs.Hit().T != 0 {
				c.WritePixel(x, y, color)
			}
		}
	}

	canvas.Export(c, "sphere.png")
}
