package shapes

import (
	"math"
	"reflect"
	"testing"

	"github.com/adriffaud/ray-tracer-challenge/pkg/primitives"
)

func TestSphereprimitivesIntersection(t *testing.T) {
	tests := []struct {
		expected []float64
		s        Shape
		r        primitives.Ray
	}{
		{
			r: primitives.Ray{
				Origin:    primitives.Point{Z: -5},
				Direction: primitives.Vector{Z: 1},
			},
			s:        Sphere(),
			expected: []float64{4.0, 6.0},
		},
		{
			r: primitives.Ray{
				Origin:    primitives.Point{Y: 1, Z: -5},
				Direction: primitives.Vector{Z: 1},
			},
			s:        Sphere(),
			expected: []float64{5.0, 5.0},
		},
		{
			r: primitives.Ray{
				Origin:    primitives.Point{Y: 2, Z: -5},
				Direction: primitives.Vector{Z: 1},
			},
			s:        Sphere(),
			expected: []float64{},
		},
		{
			r: primitives.Ray{
				Origin:    primitives.Point{},
				Direction: primitives.Vector{Z: 1},
			},
			s:        Sphere(),
			expected: []float64{-1.0, 1.0},
		},
		{
			r: primitives.Ray{
				Origin:    primitives.Point{Z: 5},
				Direction: primitives.Vector{Z: 1},
			},
			s:        Sphere(),
			expected: []float64{-6.0, -4.0},
		},
	}

	for _, test := range tests {
		xs := test.s.Intersect(test.r)

		if len(xs) != len(test.expected) {
			t.Fatalf("expected %d intersections. got=%d", len(test.expected), len(xs))
		}
		if len(xs) < 1 {
			continue
		}

		if xs[0].T != test.expected[0] {
			t.Fatalf("expected first intersection to be at %f. got=%f", test.expected[0], xs[0].T)
		}
		if xs[1].T != test.expected[1] {
			t.Fatalf("expected second intersection to be at %f. got=%f", test.expected[1], xs[1].T)
		}
	}
}

func TestSphereDefaultTransformation(t *testing.T) {
	s := Sphere()
	expected := primitives.IdentityMatrix()

	if !reflect.DeepEqual(s.Transform, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, s.Transform)
	}
}

func TestScaledSphereRayIntersection(t *testing.T) {
	r := primitives.Ray{
		Origin:    primitives.Point{Z: -5},
		Direction: primitives.Vector{Z: 1},
	}
	s := Sphere()
	s.Transform = primitives.Scaling(2, 2, 2)
	xs := s.Intersect(r)

	if len(xs) != 2 {
		t.Fatalf("expected 2 values. got=%d", len(xs))
	}
	if xs[0].T != 3 {
		t.Fatalf("expected 3. got=%f", xs[0].T)
	}
	if xs[1].T != 7 {
		t.Fatalf("expected 7. got=%f", xs[1].T)
	}
}

func TestTranslatedSphereRayIntersection(t *testing.T) {
	r := primitives.Ray{
		Origin:    primitives.Point{Z: -5},
		Direction: primitives.Vector{Z: 1},
	}
	s := Sphere()
	s.Transform = primitives.Translation(5, 0, 0)
	xs := s.Intersect(r)

	if len(xs) != 0 {
		t.Fatalf("expected no values. got=%d", len(xs))
	}
}

func TestSphereNormal(t *testing.T) {
	s := Sphere()
	tests := []struct {
		n        primitives.Vector
		expected primitives.Vector
	}{
		{s.NormalAt(primitives.Point{X: 1}), primitives.Vector{X: 1}},
		{s.NormalAt(primitives.Point{Y: 1}), primitives.Vector{Y: 1}},
		{s.NormalAt(primitives.Point{Z: 1}), primitives.Vector{Z: 1}},
		{
			s.NormalAt(primitives.Point{X: math.Sqrt(3) / 3, Y: math.Sqrt(3) / 3, Z: math.Sqrt(3) / 3}),
			primitives.Vector{X: math.Sqrt(3) / 3, Y: math.Sqrt(3) / 3, Z: math.Sqrt(3) / 3},
		},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(test.n, test.expected) {
			t.Fatalf("expected %+v. got=%+v", test.expected, test.n)
		}
	}
}

func TestSphereNormalIsNormalized(t *testing.T) {
	s := Sphere()
	n := s.NormalAt(primitives.Point{X: math.Sqrt(3) / 3, Y: math.Sqrt(3) / 3, Z: math.Sqrt(3) / 3})
	expected := n.Normalize()

	if !reflect.DeepEqual(n, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, n)
	}
}
