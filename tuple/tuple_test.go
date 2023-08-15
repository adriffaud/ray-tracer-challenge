package tuple

import (
	"reflect"
	"testing"
)

func TestTupleIsPoint(t *testing.T) {
	a := NewPoint(4.3, -4.2, 3.1)

	if reflect.TypeOf(a) != reflect.TypeOf(&Point{}) {
		t.Fatalf("a is not a Point. got=%T", a)
	}

	if a.X() != 4.3 {
		t.Fatalf("x is not 4.3. got=%f", a.X())
	}

	if a.Y() != -4.2 {
		t.Fatalf("y is not -4.2. got=%f", a.Y())
	}

	if a.Z() != 3.1 {
		t.Fatalf("z is not 3.1. got=%f", a.Z())
	}

	if a.W() != 1.0 {
		t.Fatalf("w is not 1.0. got=%f", a.W())
	}
}

func TestTupleIsVector(t *testing.T) {
	a := NewVector(4.3, -4.2, 3.1)

	if reflect.TypeOf(a) != reflect.TypeOf(&Vector{}) {
		t.Fatalf("a is not a Vector. got=%T", a)
	}

	if a.X() != 4.3 {
		t.Fatalf("x is not 4.3. got=%f", a.X())
	}

	if a.Y() != -4.2 {
		t.Fatalf("y is not -4.2. got=%f", a.Y())
	}

	if a.Z() != 3.1 {
		t.Fatalf("z is not 3.1. got=%f", a.Z())
	}

	if a.W() != 0.0 {
		t.Fatalf("w is not 0.0. got=%f", a.W())
	}
}

func TestTupleAddition(t *testing.T) {
	point := NewPoint(3, -2, 5)
	vector := NewVector(-2, 3, 1)
	expected := NewPoint(1, 1, 6)

	actual, err := Add(point, vector)
	if err != nil {
		t.Fatal("Expected no error.")
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected a Point{1, 1, 6}. got=%v", actual)
	}
}

func TestVectorAddition(t *testing.T) {
	vector1 := NewVector(3, -2, 5)
	vector2 := NewVector(-2, 3, 1)
	expected := NewVector(1, 1, 6)

	actual, err := Add(vector1, vector2)
	if err != nil {
		t.Fatal("Expected no error.")
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected a Point{1, 1, 6}. got=%v", actual)
	}
}

func TestPointAddition(t *testing.T) {
	point := NewPoint(3, -2, 5)
	vector := NewPoint(-2, 3, 1)

	_, err := Add(point, vector)
	if err == nil {
		t.Fatal("Expected an error.")
	}
}
