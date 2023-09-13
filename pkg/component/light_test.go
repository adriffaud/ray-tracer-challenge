package component

import (
	"math"
	"reflect"
	"testing"

	"github.com/adriffaud/ray-tracer-challenge/pkg/color"
	"github.com/adriffaud/ray-tracer-challenge/pkg/float"
	"github.com/adriffaud/ray-tracer-challenge/pkg/primitives"
)

func TestPointLight(t *testing.T) {
	l := Light{
		Intensity: color.Color{R: 1, G: 1, B: 1},
	}

	if !reflect.DeepEqual(l.Intensity, color.Color{R: 1, G: 1, B: 1}) {
		t.Fail()
	}

	if !reflect.DeepEqual(l.Position, primitives.Point{}) {
		t.Fail()
	}
}

func TestLighting(t *testing.T) {
	m := primitives.NewMaterial()
	position := primitives.Point{}

	tests := []struct {
		eye, normal primitives.Vector
		light       Light
		expected    color.Color
	}{
		{
			primitives.Vector{Z: -1},
			primitives.Vector{Z: -1},
			Light{
				Position:  primitives.Point{Z: -10},
				Intensity: color.Color{R: 1, G: 1, B: 1},
			},
			color.Color{R: 1.9, G: 1.9, B: 1.9},
		},
		{
			primitives.Vector{Y: math.Sqrt(2) / 2, Z: -math.Sqrt(2) / 2},
			primitives.Vector{Z: -1},
			Light{
				Position:  primitives.Point{Z: -10},
				Intensity: color.Color{R: 1, G: 1, B: 1},
			},
			color.Color{R: 1, G: 1, B: 1},
		},
		{
			primitives.Vector{Z: -1},
			primitives.Vector{Z: -1},
			Light{
				Position:  primitives.Point{Y: 10, Z: -10},
				Intensity: color.Color{R: 1, G: 1, B: 1},
			},
			color.Color{R: 0.7364, G: 0.7364, B: 0.7364},
		},
		{
			primitives.Vector{Y: -math.Sqrt(2) / 2, Z: -math.Sqrt(2) / 2},
			primitives.Vector{Z: -1},
			Light{
				Position:  primitives.Point{Y: 10, Z: -10},
				Intensity: color.Color{R: 1, G: 1, B: 1},
			},
			color.Color{R: 1.6364, G: 1.6364, B: 1.6364},
		},
		{
			primitives.Vector{Z: -1},
			primitives.Vector{Z: -1},
			Light{
				Position:  primitives.Point{Z: 10},
				Intensity: color.Color{R: 1, G: 1, B: 1},
			},
			color.Color{R: 0.1, G: 0.1, B: 0.1},
		},
	}

	for _, test := range tests {
		res := Lighting(m, test.light, position, test.eye, test.normal)

		if !float.ApproxEq(res.R, test.expected.R) {
			t.Fatalf("expected %f. got=%f", test.expected.R, res.R)
		}
		if !float.ApproxEq(res.G, test.expected.G) {
			t.Fatalf("expected %f. got=%f", test.expected.G, res.G)
		}
		if !float.ApproxEq(res.B, test.expected.B) {
			t.Fatalf("expected %f. got=%f", test.expected.B, res.B)
		}
	}
}
