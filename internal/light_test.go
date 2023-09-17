package internal

import (
	"math"
	"reflect"
	"testing"
)

func TestPointLight(t *testing.T) {
	l := Light{
		Intensity: Color{R: 1, G: 1, B: 1},
	}

	if !reflect.DeepEqual(l.Intensity, Color{R: 1, G: 1, B: 1}) {
		t.Fail()
	}

	if !reflect.DeepEqual(l.Position, Point{}) {
		t.Fail()
	}
}

func TestLighting(t *testing.T) {
	m := NewMaterial()
	position := Point{}

	tests := []struct {
		eye, normal Vector
		light       Light
		expected    Color
	}{
		{
			Vector{Z: -1},
			Vector{Z: -1},
			Light{
				Position:  Point{Z: -10},
				Intensity: Color{R: 1, G: 1, B: 1},
			},
			Color{R: 1.9, G: 1.9, B: 1.9},
		},
		{
			Vector{Y: math.Sqrt(2) / 2, Z: -math.Sqrt(2) / 2},
			Vector{Z: -1},
			Light{
				Position:  Point{Z: -10},
				Intensity: Color{R: 1, G: 1, B: 1},
			},
			Color{R: 1, G: 1, B: 1},
		},
		{
			Vector{Z: -1},
			Vector{Z: -1},
			Light{
				Position:  Point{Y: 10, Z: -10},
				Intensity: Color{R: 1, G: 1, B: 1},
			},
			Color{R: 0.7364, G: 0.7364, B: 0.7364},
		},
		{
			Vector{Y: -math.Sqrt(2) / 2, Z: -math.Sqrt(2) / 2},
			Vector{Z: -1},
			Light{
				Position:  Point{Y: 10, Z: -10},
				Intensity: Color{R: 1, G: 1, B: 1},
			},
			Color{R: 1.6364, G: 1.6364, B: 1.6364},
		},
		{
			Vector{Z: -1},
			Vector{Z: -1},
			Light{
				Position:  Point{Z: 10},
				Intensity: Color{R: 1, G: 1, B: 1},
			},
			Color{R: 0.1, G: 0.1, B: 0.1},
		},
	}

	for _, test := range tests {
		res := Lighting(m, test.light, position, test.eye, test.normal)
		assertColorEquals(t, res, test.expected)
	}
}
