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
		inShadow    bool
	}{
		{
			Vector{Z: -1},
			Vector{Z: -1},
			Light{Color{1, 1, 1}, Point{Z: -10}},
			Color{1.9, 1.9, 1.9},
			false,
		},
		{
			Vector{Y: math.Sqrt(2) / 2, Z: -math.Sqrt(2) / 2},
			Vector{Z: -1},
			Light{Color{1, 1, 1}, Point{Z: -10}},
			Color{1, 1, 1},
			false,
		},
		{
			Vector{Z: -1},
			Vector{Z: -1},
			Light{Color{1, 1, 1}, Point{Y: 10, Z: -10}},
			Color{0.7364, 0.7364, 0.7364},
			false,
		},
		{
			Vector{Y: -math.Sqrt(2) / 2, Z: -math.Sqrt(2) / 2},
			Vector{Z: -1},
			Light{Color{1, 1, 1}, Point{Y: 10, Z: -10}},
			Color{1.6364, 1.6364, 1.6364},
			false,
		},
		{
			Vector{Z: -1},
			Vector{Z: -1},
			Light{Color{1, 1, 1}, Point{Z: 10}},
			Color{0.1, 0.1, 0.1},
			false,
		},
		{
			Vector{Z: -1},
			Vector{Z: -1},
			Light{Color{1, 1, 1}, Point{Z: -10}},
			Color{0.1, 0.1, 0.1},
			true,
		},
	}

	for _, test := range tests {
		res := Lighting(m, test.light, position, test.eye, test.normal, test.inShadow)
		assertColorEquals(t, test.expected, res)
	}
}
