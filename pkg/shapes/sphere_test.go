package shapes

import (
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
				Origin:    primitives.Point{X: 0, Y: 0, Z: -5},
				Direction: primitives.Vector{X: 0, Y: 0, Z: 1},
			},
			s:        Sphere(),
			expected: []float64{4.0, 6.0},
		},
		{
			r: primitives.Ray{
				Origin:    primitives.Point{X: 0, Y: 1, Z: -5},
				Direction: primitives.Vector{X: 0, Y: 0, Z: 1},
			},
			s:        Sphere(),
			expected: []float64{5.0, 5.0},
		},
		{
			r: primitives.Ray{
				Origin:    primitives.Point{X: 0, Y: 2, Z: -5},
				Direction: primitives.Vector{X: 0, Y: 0, Z: 1},
			},
			s:        Sphere(),
			expected: []float64{},
		},
		{
			r: primitives.Ray{
				Origin:    primitives.Point{X: 0, Y: 0, Z: 0},
				Direction: primitives.Vector{X: 0, Y: 0, Z: 1},
			},
			s:        Sphere(),
			expected: []float64{-1.0, 1.0},
		},
		{
			r: primitives.Ray{
				Origin:    primitives.Point{X: 0, Y: 0, Z: 5},
				Direction: primitives.Vector{X: 0, Y: 0, Z: 1},
			},
			s:        Sphere(),
			expected: []float64{-6.0, -4.0},
		},
	}

	for _, test := range tests {
		xs := test.s.Intersect(test.r)

		if len(xs) < len(test.expected) {
			t.Fatalf("expected %d intersections. got=%d", len(test.expected), len(xs))
		}
		if len(xs) < 1 {
			continue
		}

		if xs[0].Object != test.s {
			t.Fatalf("expected first intersection to be at %f. got=%f", test.expected[0], xs[0])
		}
		if xs[1].Object != test.s {
			t.Fatalf("expected second intersection to be at %f. got=%f", test.expected[1], xs[1])
		}
	}
}
