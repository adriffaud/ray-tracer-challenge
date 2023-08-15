package color

import (
	"reflect"
	"testing"

	"github.com/adriffaud/ray-tracer-challenge/float"
)

func assertEquals(t *testing.T, c1, c2 Color) {
	if !float.ApproxEq(c1.Red, c2.Red) {
		t.Fatalf("expected %+v. got=%+v", c1, c2)
	}
	if !float.ApproxEq(c1.Green, c2.Green) {
		t.Fatalf("expected %+v. got=%+v", c1, c2)
	}
	if !float.ApproxEq(c1.Blue, c2.Blue) {
		t.Fatalf("expected %+v. got=%+v", c1, c2)
	}
}

func TestColorRGBTuples(t *testing.T) {
	c := NewColor(-0.5, 0.4, 1.7)

	if c.Red != -0.5 {
		t.Fatalf("expected -0.5. got=%f", c.Red)
	}

	if c.Green != 0.4 {
		t.Fatalf("expected 0.4. got=%f", c.Green)
	}

	if c.Blue != 1.7 {
		t.Fatalf("expected 1.7. got=%f", c.Blue)
	}
}

func TestColorAddition(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	expected := NewColor(1.6, 0.7, 1.0)

	actual := Add(c1, c2)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", actual, expected)
	}
}

func TestColorSubtraction(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	expected := NewColor(0.2, 0.5, 0.5)

	actual := Sub(c1, c2)
	assertEquals(t, *actual, *expected)
}

func TestColorScalarMultiplication(t *testing.T) {
	c := NewColor(0.2, 0.3, 0.4)
	expected := NewColor(0.4, 0.6, 0.8)

	actual := MultiplyByScalar(c, 2)
	assertEquals(t, *actual, *expected)
}

func TestColorMultiplication(t *testing.T) {
	c1 := NewColor(1, 0.2, 0.4)
	c2 := NewColor(0.9, 1, 0.1)
	expected := NewColor(0.9, 0.2, 0.04)

	actual := Multiply(c1, c2)
	assertEquals(t, *actual, *expected)
}
