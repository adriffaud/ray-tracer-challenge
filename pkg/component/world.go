package component

import (
	"sort"

	"github.com/adriffaud/ray-tracer-challenge/pkg/color"
	"github.com/adriffaud/ray-tracer-challenge/pkg/primitives"
	"github.com/adriffaud/ray-tracer-challenge/pkg/shape"
)

type World struct {
	Objects []shape.Shape
	Light   Light
}

func NewWorld() World {
	light := Light{
		Position:  primitives.Point{X: -10, Y: 10, Z: -10},
		Intensity: color.Color{R: 1, G: 1, B: 1},
	}

	s1 := shape.Sphere()
	s1.Material.Color = color.Color{R: 0.8, G: 1, B: 0.6}
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2

	s2 := shape.Sphere()
	s2.Transform = primitives.Scaling(0.5, 0.5, 0.5)

	return World{Light: light, Objects: []shape.Shape{s1, s2}}
}

func (w World) Intersect(r primitives.Ray) shape.Intersections {
	var i shape.Intersections
	for _, o := range w.Objects {
		i = append(i, o.Intersect(r)...)
	}
	sort.Sort(i)

	return i
}
