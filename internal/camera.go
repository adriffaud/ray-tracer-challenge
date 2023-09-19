package internal

import (
	"math"
)

type Camera struct {
	Transform                                     Matrix
	Width, Height                                 int
	FieldOfView, HalfWidth, HalfHeight, PixelSize float64
}

func NewCamera(width, height int, fieldOfView float64) Camera {
	halfView := math.Tan(fieldOfView / 2.0)
	aspect := float64(width) / float64(height)

	c := Camera{
		Width:       width,
		Height:      height,
		FieldOfView: fieldOfView,
		Transform:   IdentityMatrix(),
	}

	if aspect >= 1 {
		c.HalfWidth = halfView
		c.HalfHeight = halfView / aspect
	} else {
		c.HalfWidth = halfView * aspect
		c.HalfHeight = halfView
	}
	c.PixelSize = (c.HalfWidth * 2) / float64(width)

	return c
}

func (c Camera) RayForPixel(x, y int) Ray {
	xOffset := (float64(x) + 0.5) * c.PixelSize
	yOffset := (float64(y) + 0.5) * c.PixelSize

	worldX := c.HalfWidth - xOffset
	worldY := c.HalfHeight - yOffset

	inv, err := c.Transform.Inverse()
	if err != nil {
		panic(err)
	}
	pixel := Point{X: worldX, Y: worldY, Z: -1}.MultiplyMatrix(inv)
	origin := Point{}.MultiplyMatrix(inv)
	direction := pixel.SubPoint(origin).Normalize()

	return Ray{Origin: origin, Direction: direction}
}

func (c Camera) Render(w World) Canvas {
	image := NewCanvas(c.Width, c.Height)

	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			r := c.RayForPixel(x, y)
			color := w.ColorAt(r)
			image.WritePixel(x, y, color)
		}
	}

	return *image
}
