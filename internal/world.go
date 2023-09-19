package internal

import "sort"

type World struct {
	Objects []Shape
	Lights  []Light
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

	return World{Lights: []Light{light}, Objects: []Shape{s1, s2}}
}

func (w World) Intersect(r Ray) Intersections {
	var i Intersections
	for _, o := range w.Objects {
		i = append(i, o.Intersect(r)...)
	}
	sort.Sort(i)
	return i
}

func (w World) ShadeHit(comps Computations, light Light) Color {
	return Lighting(comps.Object.Material, light, comps.Point, comps.EyeV, comps.NormalV)
}

func (w World) ColorAt(r Ray) Color {
	result := Color{}
	intersections := w.Intersect(r)
	intersection, hit := intersections.Hit()
	if !hit {
		return result
	}

	comps := intersection.PrepareComputations(r)
	for _, light := range w.Lights {
		result = result.Add(w.ShadeHit(comps, light))
	}

	return result
}
