package internal

import "sort"

type World struct {
	Objects []Shape
	Light   Light
}

func NewWorld() World {
	light := Light{
		Position:  Point{X: -10, Y: 10, Z: -10},
		Intensity: Color{R: 1, G: 1, B: 1},
	}

	s1 := Sphere()
	s1.Material.Color = Color{R: 0.8, G: 1, B: 0.6}
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2

	s2 := Sphere()
	s2.Transform = Scaling(0.5, 0.5, 0.5)

	return World{Light: light, Objects: []Shape{s1, s2}}
}

func (w World) Intersect(r Ray) Intersections {
	var i Intersections
	for _, o := range w.Objects {
		i = append(i, o.Intersect(r)...)
	}
	sort.Sort(i)

	return i
}

func (w World) ShadeHit(comps Computations) Color {
	return Lighting(comps.Object.Material, w.Light, comps.Point, comps.EyeV, comps.NormalV)
}

func (w World) ColorAt(r Ray) Color {
	xs := w.Intersect(r)
	i, hit := xs.Hit()
	if !hit {
		return Color{}
	}
	comps := i.PrepareComputations(r)
	return w.ShadeHit(comps)
}
