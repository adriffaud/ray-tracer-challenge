package main

import (
	"github.com/adriffaud/ray-tracer-challenge/pkg/canvas"
	"github.com/adriffaud/ray-tracer-challenge/pkg/color"
	"github.com/adriffaud/ray-tracer-challenge/pkg/component"
	"github.com/adriffaud/ray-tracer-challenge/pkg/primitives"
	"github.com/adriffaud/ray-tracer-challenge/pkg/shape"
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
	shape := shape.Sphere()
	shape.Material = primitives.NewMaterial()
	shape.Material.Color = color.Color{R: 1, G: 0.2, B: 1}

	light := component.Light{
		Position:  primitives.Point{X: -10, Y: 10, Z: -10},
		Intensity: color.Color{R: 1, G: 1, B: 1},
	}

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
			hit := xs.Hit()

			if hit.T != 0 {
				p := r.Position(hit.T)
				normal := hit.Object.NormalAt(p)
				eye := r.Direction.Negate()
				col := component.Lighting(hit.Object.Material, light, p, eye, normal)

				c.WritePixel(x, y, col)
			}
		}
	}

	canvas.Export(c, "sphere.png")
}
