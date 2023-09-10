package component

import (
	"math"
	"reflect"
	"testing"

	"github.com/adriffaud/ray-tracer-challenge/pkg/color"
	"github.com/adriffaud/ray-tracer-challenge/pkg/float"
	"github.com/adriffaud/ray-tracer-challenge/pkg/primitives"
)

func TestDefaultMaterial(t *testing.T) {
	m := NewMaterial()
	color := color.Color{R: 1, G: 1, B: 1}

	if !reflect.DeepEqual(m.Color, color) {
		t.Fatalf("expected %+v. got=%+v", color, m.Color)
	}
	if m.Ambient != 0.1 {
		t.Fatalf("expected 0.1. got=%f", m.Ambient)
	}
	if m.Diffuse != 0.9 {
		t.Fatalf("expected 0.9. got=%f", m.Diffuse)
	}
	if m.Specular != 0.9 {
		t.Fatalf("expected 0.9. got=%f", m.Specular)
	}
	if m.Shininess != 200.0 {
		t.Fatalf("expected 200. got=%f", m.Shininess)
	}
}

func TestLighting(t *testing.T) {
	m := NewMaterial()
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
