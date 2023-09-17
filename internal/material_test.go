package internal

import (
	"reflect"
	"testing"
)

func TestDefaultMaterial(t *testing.T) {
	m := NewMaterial()
	color := Color{R: 1, G: 1, B: 1}

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
