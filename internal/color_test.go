package internal

import (
	"reflect"
	"testing"
)

func assertColorEquals(t *testing.T, expected, actual Color) {
	if !ApproxEq(expected.R, actual.R) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
	if !ApproxEq(expected.G, actual.G) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
	if !ApproxEq(expected.B, actual.B) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
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

	actual := c1.Add(c2)
	assertDeepEqual(t, actual, expected)
}

func TestColorSubtraction(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}
	expected := Color{0.2, 0.5, 0.5}

	actual := c1.Sub(c2)
	assertColorEquals(t, actual, expected)
}

func TestColorScalarMultiplication(t *testing.T) {
	c := Color{R: 0.2, G: 0.3, B: 0.4}
	expected := Color{R: 0.4, G: 0.6, B: 0.8}

	actual := c.MultiplyScalar(2)
	assertColorEquals(t, actual, expected)
}

func TestColorMultiplication(t *testing.T) {
	c1 := Color{R: 1, G: 0.2, B: 0.4}
	c2 := Color{R: 0.9, G: 1, B: 0.1}
	expected := Color{R: 0.9, G: 0.2, B: 0.04}

	actual := c1.Multiply(c2)
	assertColorEquals(t, actual, expected)
}
