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
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Point{}) {
		t.Fatalf("result is not a Point. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected a Point{1, 1, 6}. got=%v", actual)
	}
}

func TestVectorAddition(t *testing.T) {
	vector1 := NewVector(3, -2, 5)
	vector2 := NewVector(-2, 3, 1)
	expected := NewVector(1, 1, 6)

	actual, err := Add(vector1, vector2)
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Vector{}) {
		t.Fatalf("result is not a Vector. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected a Point{1, 1, 6}. got=%v", actual)
	}
}

func TestPointAddition(t *testing.T) {
	point1 := NewPoint(3, -2, 5)
	point2 := NewPoint(-2, 3, 1)

	_, err := Add(point1, point2)
	if err == nil {
		t.Fatal("expected an error.")
	}
}

func TestPointSubtraction(t *testing.T) {
	point1 := NewPoint(3, 2, 1)
	point2 := NewPoint(5, 6, 7)
	expected := NewVector(-2, -4, -6)

	actual, err := Sub(point1, point2)
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Vector{}) {
		t.Fatalf("result is not a Vector. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected a Vector{-2, -4, -6}. got=%v", actual)
	}
}

func TestVectorFromPointSubtraction(t *testing.T) {
	point := NewPoint(3, 2, 1)
	vector := NewVector(5, 6, 7)
	expected := NewPoint(-2, -4, -6)

	actual, err := Sub(point, vector)
	if err != nil {
		t.Fatal("Expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Point{}) {
		t.Fatalf("result is not a Point. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected a Point{-2, -4, -6}. got=%v", actual)
	}
}

func TestVectorSubtraction(t *testing.T) {
	vector1 := NewVector(3, 2, 1)
	vector2 := NewVector(5, 6, 7)
	expected := NewVector(-2, -4, -6)

	actual, err := Sub(vector1, vector2)
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Vector{}) {
		t.Fatalf("result is not a Vector. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected a Vector{-2, -4, -6}. got=%v", actual)
	}
}

func TestPointFromVectorSubtraction(t *testing.T) {
	vector := NewVector(3, 2, 1)
	point := NewPoint(5, 6, 7)

	_, err := Sub(vector, point)
	if err == nil {
		t.Fatal("expected an error.")
	}
}

func TestTupleNegationWithSub(t *testing.T) {
	zero := NewZeroVector()
	vector := NewVector(1, -2, 3)
	expected := NewVector(-1, 2, -3)

	actual, err := Sub(zero, vector)
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Vector{}) {
		t.Fatalf("result is not a Vector. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected a Vector{-1, 2, -3}. got=%v", actual)
	}
}

func TestVectorNegation(t *testing.T) {
	vector := NewVector(1, -2, 3)
	expected := NewVector(-1, 2, -3)

	actual, err := Negate(vector)
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Vector{}) {
		t.Fatalf("result is not a Vector. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected a Vector{-1, 2, -3}. got=%v", actual)
	}
}

func TestPointNegation(t *testing.T) {
	point := NewPoint(1, -2, 3)
	expected := NewPoint(-1, 2, -3)

	actual, err := Negate(point)
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Point{}) {
		t.Fatalf("result is not a Point. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected a Point{-1, 2, -3}. got=%v", actual)
	}
}
