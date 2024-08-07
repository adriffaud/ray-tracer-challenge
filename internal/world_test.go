package internal

import (
	"reflect"
	"slices"
	"testing"
)

func TestDefaultWorld(t *testing.T) {
	light := Light{Color{1, 1, 1}, Point{-10, 10, -10}}

	s1 := Sphere()
	s1.Material.Color = Color{R: 0.8, G: 1, B: 0.6}
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2

	s2 := Sphere()
	s2.Transform = Scaling(0.5, 0.5, 0.5)

	w := NewWorld()

	assertDeepEqual(t, w.Lights[0], light)

	if !slices.ContainsFunc(w.Objects, func(s Shape) bool {
		return reflect.DeepEqual(s, s1)
	}) {
		t.Fatalf("expected to contain %+v", s1)
	}

	if !slices.ContainsFunc(w.Objects, func(s Shape) bool {
		return reflect.DeepEqual(s, s2)
	}) {
		t.Fatalf("expected to contain %+v", s2)
	}
}

func TestWorldIntersect(t *testing.T) {
	w := NewWorld()
	r := Ray{Point{Z: -5}, Vector{Z: 1}}
	xs := w.Intersect(r)

	if len(xs) != 4 {
		t.Fatalf("expected 4 items. got=%d", len(xs))
	}
	if xs[0].Distance != 4 {
		t.Fail()
	}
	if xs[1].Distance != 4.5 {
		t.Fail()
	}
	if xs[2].Distance != 5.5 {
		t.Fail()
	}
	if xs[3].Distance != 6 {
		t.Fail()
	}
}

func TestIntersectionShading(t *testing.T) {
	w := NewWorld()
	r := Ray{Point{Z: -5}, Vector{Z: 1}}
	s := w.Objects[0]
	i := Intersection{s, 4}
	comps := i.PrepareComputations(r)
	c := w.ShadeHit(comps, w.Lights[0])
	expected := Color{0.38066, 0.47583, 0.2855}
	assertColorEquals(t, expected, c)
}

func TestInsideIntersectionShading(t *testing.T) {
	w := NewWorld()
	w.Lights = []Light{{Color{R: 1, G: 1, B: 1}, Point{Y: 0.25}}}
	r := Ray{Point{}, Vector{Z: 1}}
	s := w.Objects[1]
	i := Intersection{s, 0.5}
	comps := i.PrepareComputations(r)
	c := w.ShadeHit(comps, w.Lights[0])
	expected := Color{0.90498, 0.90498, 0.90498}
	assertColorEquals(t, expected, c)
}

func TestColors(t *testing.T) {
	tests := []struct {
		w        World
		r        Ray
		expected Color
	}{
		{NewWorld(), Ray{Point{Z: -5}, Vector{Y: 1}}, Color{}},
		{NewWorld(), Ray{Point{Z: -5}, Vector{Z: 1}}, Color{0.38066, 0.47583, 0.2855}},
	}

	for _, test := range tests {
		c := test.w.ColorAt(test.r)
		assertColorEquals(t, test.expected, c)
	}
}

func TestBehindRayIntersectionColor(t *testing.T) {
	w := NewWorld()
	w.Objects[0].Material.Ambient = 1
	w.Objects[1].Material.Ambient = 1
	r := Ray{Point{Z: 0.75}, Vector{Z: -1}}
	c := w.ColorAt(r)
	inner := w.Objects[1]
	assertColorEquals(t, inner.Material.Color, c)
}

func TestShadow(t *testing.T) {
	w := NewWorld()

	tests := []struct {
		p          Point
		isShadowed bool
	}{
		{Point{Y: 10}, false},
		{Point{10, -10, 10}, true},
		{Point{-20, 20, -20}, false},
		{Point{-2, 2, -2}, false},
	}

	for _, test := range tests {
		actual := w.IsShadowed(test.p)

		if actual != test.isShadowed {
			t.Fatalf("expected %t. got=%t", test.isShadowed, actual)
		}
	}
}

func TestShadowIntersection(t *testing.T) {
	w := World{
		[]Shape{Sphere()},
		[]Light{{Color{1, 1, 1}, Point{Z: -10}}},
	}

	s2 := Sphere()
	s2.Transform = Translation(0, 0, 10)

	w.Objects = append(w.Objects, s2)

	r := Ray{Point{Z: 5}, Vector{Z: 1}}
	i := Intersection{s2, 4}
	comps := i.PrepareComputations(r)
	c := w.ShadeHit(comps, w.Lights[0])
	expected := Color{0.1, 0.1, 0.1}

	assertColorEquals(t, expected, c)
}
