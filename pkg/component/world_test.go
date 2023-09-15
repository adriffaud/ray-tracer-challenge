package component

import (
	"reflect"
	"testing"

	"slices"

	"github.com/adriffaud/ray-tracer-challenge/pkg/color"
	"github.com/adriffaud/ray-tracer-challenge/pkg/float"
	"github.com/adriffaud/ray-tracer-challenge/pkg/primitives"
	"github.com/adriffaud/ray-tracer-challenge/pkg/shape"
)

func TestDefaultWorld(t *testing.T) {
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

	w := NewWorld()

	if !reflect.DeepEqual(w.Light, light) {
		t.Fatalf("expected %+v. got=%+v", light, w.Light)
	}

	if !slices.ContainsFunc(w.Objects, func(s shape.Shape) bool {
		return reflect.DeepEqual(s, s1)
	}) {
		t.Fatalf("expected to contain %+v", s1)
	}

	if !slices.ContainsFunc(w.Objects, func(s shape.Shape) bool {
		return reflect.DeepEqual(s, s2)
	}) {
		t.Fatalf("expected to contain %+v", s2)
	}
}

func TestWorldIntersect(t *testing.T) {
	w := NewWorld()
	r := primitives.Ray{
		Origin:    primitives.Point{Z: -5},
		Direction: primitives.Vector{Z: 1},
	}
	xs := w.Intersect(r)

	if len(xs) != 4 {
		t.Fatalf("expected 4 items. got=%d", len(xs))
	}
	if xs[0].Distance != 4 {
		t.Fail()
	}
}

func TestIntersectionShading(t *testing.T) {
	w := NewWorld()
	r := primitives.Ray{
		Origin:    primitives.Point{Z: -5},
		Direction: primitives.Vector{Z: 1},
	}
	s := w.Objects[0]
	i := shape.Intersection{Object: s, Distance: 4}
	comps := i.PrepareComputations(r)
	c := w.ShadeHit(comps)
	expected := color.Color{R: 0.38066, G: 0.47583, B: 0.2855}

	if !float.ApproxEq(c.R, expected.R) {
		t.Fatalf("expected %+v. got=%+v", c, expected)
	}
	if !float.ApproxEq(c.G, expected.G) {
		t.Fatalf("expected %+v. got=%+v", c, expected)
	}
	if !float.ApproxEq(c.B, expected.B) {
		t.Fatalf("expected %+v. got=%+v", c, expected)
	}
}

func TestInsideIntersectionShading(t *testing.T) {
	w := NewWorld()
	w.Light = Light{
		Position:  primitives.Point{Y: 0.25},
		Intensity: color.Color{R: 1, G: 1, B: 1},
	}
	r := primitives.Ray{
		Origin:    primitives.Point{},
		Direction: primitives.Vector{Z: 1},
	}
	s := w.Objects[1]
	i := shape.Intersection{Object: s, Distance: 0.5}
	comps := i.PrepareComputations(r)
	c := w.ShadeHit(comps)
	expected := color.Color{R: 0.90498, G: 0.90498, B: 0.90498}

	if !float.ApproxEq(c.R, expected.R) {
		t.Fatalf("expected %+v. got=%+v", c, expected)
	}
	if !float.ApproxEq(c.G, expected.G) {
		t.Fatalf("expected %+v. got=%+v", c, expected)
	}
	if !float.ApproxEq(c.B, expected.B) {
		t.Fatalf("expected %+v. got=%+v", c, expected)
	}
}

func TestColors(t *testing.T) {
	tests := []struct {
		w        World
		r        primitives.Ray
		expected color.Color
	}{
		{
			NewWorld(),
			primitives.Ray{
				Origin:    primitives.Point{Z: -5},
				Direction: primitives.Vector{Y: 1},
			},
			color.Color{},
		},
		{
			NewWorld(),
			primitives.Ray{
				Origin:    primitives.Point{Z: -5},
				Direction: primitives.Vector{Z: 1},
			},
			color.Color{R: 0.38066, G: 0.47583, B: 0.2855},
		},
	}

	for _, test := range tests {
		c := test.w.ColorAt(test.r)

		if !float.ApproxEq(c.R, test.expected.R) {
			t.Fatalf("expected %+v. got=%+v", test.expected, c)
		}
		if !float.ApproxEq(c.G, test.expected.G) {
			t.Fatalf("expected %+v. got=%+v", test.expected, c)
		}
		if !float.ApproxEq(c.B, test.expected.B) {
			t.Fatalf("expected %+v. got=%+v", test.expected, c)
		}
	}
}

func TestBehindRayIntersectionColor(t *testing.T) {
	w := NewWorld()
	outer := w.Objects[0]
	outer.Material.Ambient = 1
	inner := w.Objects[1]
	inner.Material.Ambient = 1
	r := primitives.Ray{
		Origin:    primitives.Point{Z: 0.75},
		Direction: primitives.Vector{Z: -1},
	}
	c := w.ColorAt(r)

	if !float.ApproxEq(c.R, inner.Material.Color.R) {
		t.Fatalf("expected %+v. got=%+v", inner.Material.Color, c)
	}
	if !float.ApproxEq(c.G, inner.Material.Color.G) {
		t.Fatalf("expected %+v. got=%+v", inner.Material.Color, c)
	}
	if !float.ApproxEq(c.B, inner.Material.Color.B) {
		t.Fatalf("expected %+v. got=%+v", inner.Material.Color, c)
	}
}
