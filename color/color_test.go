package color

import (
	"reflect"
	"testing"

	"github.com/adriffaud/ray-tracer-challenge/float"
)

func assertEquals(t *testing.T, c1, c2 Color) {
	if !float.ApproxEq(c1.R, c2.R) {
		t.Fatalf("expected %+v. got=%+v", c1, c2)
	}
	if !float.ApproxEq(c1.G, c2.G) {
		t.Fatalf("expected %+v. got=%+v", c1, c2)
	}
	if !float.ApproxEq(c1.B, c2.B) {
		t.Fatalf("expected %+v. got=%+v", c1, c2)
	}
}

func TestColorRGBTuples(t *testing.T) {
	c := Color{R: -0.5, G: 0.4, B: 1.7}

	if c.R != -0.5 {
		t.Fatalf("expected -0.5. got=%f", c.R)
	}

	if c.G != 0.4 {
		t.Fatalf("expected 0.4. got=%f", c.G)
	}

	if c.B != 1.7 {
		t.Fatalf("expected 1.7. got=%f", c.B)
	}
}

func TestColorAddition(t *testing.T) {
	c1 := Color{R: 0.9, G: 0.6, B: 0.75}
	c2 := Color{R: 0.7, G: 0.1, B: 0.25}
	expected := Color{1.6, 0.7, 1.0}

	actual := Add(c1, c2)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", actual, expected)
	}
}

func TestColorSubtraction(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}
	expected := Color{0.2, 0.5, 0.5}

	actual := Sub(c1, c2)
	assertEquals(t, actual, expected)
}

func TestColorScalarMultiplication(t *testing.T) {
	c := Color{R: 0.2, G: 0.3, B: 0.4}
	expected := Color{R: 0.4, G: 0.6, B: 0.8}

	actual := MultiplyByScalar(c, 2)
	assertEquals(t, actual, expected)
}

func TestColorMultiplication(t *testing.T) {
	c1 := Color{R: 1, G: 0.2, B: 0.4}
	c2 := Color{R: 0.9, G: 1, B: 0.1}
	expected := Color{R: 0.9, G: 0.2, B: 0.04}

	actual := Multiply(c1, c2)
	assertEquals(t, actual, expected)
}
