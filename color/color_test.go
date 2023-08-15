package color

import (
	"reflect"
	"testing"
)

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
