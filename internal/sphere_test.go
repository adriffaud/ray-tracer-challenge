package internal

import (
	"math"
	"reflect"
	"testing"
)

func TestSpherentersection(t *testing.T) {
	tests := []struct {
		expected []float64
		s        Shape
		r        Ray
	}{
		{
			r: Ray{
				Origin:    Point{Z: -5},
				Direction: Vector{Z: 1},
			},
			s:        Sphere(),
			expected: []float64{4.0, 6.0},
		},
		{
			r: Ray{
				Origin:    Point{Y: 1, Z: -5},
				Direction: Vector{Z: 1},
			},
			s:        Sphere(),
			expected: []float64{5.0, 5.0},
		},
		{
			r: Ray{
				Origin:    Point{Y: 2, Z: -5},
				Direction: Vector{Z: 1},
			},
			s:        Sphere(),
			expected: []float64{},
		},
		{
			r: Ray{
				Origin:    Point{},
				Direction: Vector{Z: 1},
			},
			s:        Sphere(),
			expected: []float64{-1.0, 1.0},
		},
		{
			r: Ray{
				Origin:    Point{Z: 5},
				Direction: Vector{Z: 1},
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

		if xs[0].Distance != test.expected[0] {
			t.Fatalf("expected first intersection to be at %f. got=%f", test.expected[0], xs[0].Distance)
		}
		if xs[1].Distance != test.expected[1] {
			t.Fatalf("expected second intersection to be at %f. got=%f", test.expected[1], xs[1].Distance)
		}
	}
}

func TestSphereDefaultTransformation(t *testing.T) {
	s := Sphere()
	expected := IdentityMatrix()

	if !reflect.DeepEqual(s.Transform, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, s.Transform)
	}
}

func TestScaledSphereRayIntersection(t *testing.T) {
	r := Ray{
		Origin:    Point{Z: -5},
		Direction: Vector{Z: 1},
	}
	s := Sphere()
	s.Transform = Scaling(2, 2, 2)
	xs := s.Intersect(r)

	if len(xs) != 2 {
		t.Fatalf("expected 2 values. got=%d", len(xs))
	}
	if xs[0].Distance != 3 {
		t.Fatalf("expected 3. got=%f", xs[0].Distance)
	}
	if xs[1].Distance != 7 {
		t.Fatalf("expected 7. got=%f", xs[1].Distance)
	}
}

func TestTranslatedSphereRayIntersection(t *testing.T) {
	r := Ray{
		Origin:    Point{Z: -5},
		Direction: Vector{Z: 1},
	}
	s := Sphere()
	s.Transform = Translation(5, 0, 0)
	xs := s.Intersect(r)

	if len(xs) != 0 {
		t.Fatalf("expected no values. got=%d", len(xs))
	}
}

func TestSphereNormal(t *testing.T) {
	s := Sphere()
	tests := []struct {
		n        Vector
		expected Vector
	}{
		{s.NormalAt(Point{X: 1}), Vector{X: 1}},
		{s.NormalAt(Point{Y: 1}), Vector{Y: 1}},
		{s.NormalAt(Point{Z: 1}), Vector{Z: 1}},
		{
			s.NormalAt(Point{X: math.Sqrt(3) / 3, Y: math.Sqrt(3) / 3, Z: math.Sqrt(3) / 3}),
			Vector{X: math.Sqrt(3) / 3, Y: math.Sqrt(3) / 3, Z: math.Sqrt(3) / 3},
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
	n := s.NormalAt(Point{X: math.Sqrt(3) / 3, Y: math.Sqrt(3) / 3, Z: math.Sqrt(3) / 3})
	expected := n.Normalize()

	if !reflect.DeepEqual(n, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, n)
	}
}

func assertVectorEquals(t *testing.T, expected, actual Vector) {
	if !ApproxEq(expected.X, actual.X) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
	if !ApproxEq(expected.Y, actual.Y) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
	if !ApproxEq(expected.Z, actual.Z) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestTranslatedSphereNormal(t *testing.T) {
	s := Sphere()
	s.Transform = Translation(0, 1, 0)
	n := s.NormalAt(Point{Y: 1.70711, Z: -0.70711})
	expected := Vector{Y: 0.70711, Z: -0.70711}

	assertVectorEquals(t, expected, n)
}

func TestTranformedSphereNormal(t *testing.T) {
	s := Sphere()
	m := Scaling(1, 0.5, 1).Multiply(RotationZ(math.Pi / 5))
	s.Transform = m
	n := s.NormalAt(Point{Y: math.Sqrt(2) / 2, Z: -math.Sqrt(2) / 2})
	expected := Vector{Y: 0.97014, Z: -0.24254}

	assertVectorEquals(t, expected, n)
}

func TestSphereDefaultMaterial(t *testing.T) {
	s := Sphere()
	expected := NewMaterial()

	if !reflect.DeepEqual(s.Material, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, s.Material)
	}
}
