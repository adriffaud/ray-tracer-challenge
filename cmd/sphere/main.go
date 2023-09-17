package main

import "github.com/adriffaud/ray-tracer-challenge/internal"

const (
	CANVAS_PIXELS = 100.0
	WALL_Z        = 10
	WALL_SIZE     = 7.0
	PIXEL_SIZE    = WALL_SIZE / CANVAS_PIXELS
	HALF          = WALL_SIZE / 2
)

func main() {
	rayOrigin := internal.Point{Z: -5}

	c := internal.NewCanvas(CANVAS_PIXELS, CANVAS_PIXELS)
	shape := internal.Sphere()
	shape.Material = internal.NewMaterial()
	shape.Material.Color = internal.Color{R: 1, G: 0.2, B: 1}

	light := internal.Light{
		Position:  internal.Point{X: -10, Y: 10, Z: -10},
		Intensity: internal.Color{R: 1, G: 1, B: 1},
	}

	for y := 0; y < CANVAS_PIXELS; y++ {
		worldY := HALF - PIXEL_SIZE*float64(y)
		for x := 0; x < CANVAS_PIXELS; x++ {
			worldX := -HALF + PIXEL_SIZE*float64(x)
			position := internal.Point{
				X: worldX,
				Y: worldY,
				Z: WALL_Z,
			}
			direction := position.SubPoint(rayOrigin).Normalize()

			r := internal.Ray{Origin: rayOrigin, Direction: direction}
			xs := shape.Intersect(r)
			i, hit := xs.Hit()

			if hit {
				p := r.Position(i.Distance)
				normal := i.Object.NormalAt(p)
				eye := r.Direction.Negate()
				col := internal.Lighting(i.Object.Material, light, p, eye, normal)

				c.WritePixel(x, y, col)
			}
		}
	}

	internal.Export(c, "sphere.png")
}
