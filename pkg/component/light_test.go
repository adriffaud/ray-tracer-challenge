package component

import (
	"reflect"
	"testing"

	"github.com/adriffaud/ray-tracer-challenge/pkg/color"
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
