package component

import (
	"reflect"
	"testing"

	"slices"

	"github.com/adriffaud/ray-tracer-challenge/pkg/color"
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
