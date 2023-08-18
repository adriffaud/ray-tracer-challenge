package tuple

import (
	"math"
	"reflect"
	"testing"
)

func TestTupleIsPoint(t *testing.T) {
	a := Point{4.3, -4.2, 3.1}

	if reflect.TypeOf(a) != reflect.TypeOf(Point{}) {
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
	a := Vector{4.3, -4.2, 3.1}

	if reflect.TypeOf(a) != reflect.TypeOf(Vector{}) {
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
	point := Point{3, -2, 5}
	vector := Vector{-2, 3, 1}
	expected := Point{1, 1, 6}

	actual, err := Add(&point, &vector)
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Point{}) {
		t.Fatalf("result is not a Point. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", &expected, actual)
	}
}

func TestVectorAddition(t *testing.T) {
	vector1 := Vector{3, -2, 5}
	vector2 := Vector{-2, 3, 1}
	expected := Vector{1, 1, 6}

	actual, err := Add(&vector1, &vector2)
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Vector{}) {
		t.Fatalf("result is not a Vector. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", &expected, actual)
	}
}

func TestPointAddition(t *testing.T) {
	point1 := Point{3, -2, 5}
	point2 := Point{-2, 3, 1}

	_, err := Add(&point1, &point2)
	if err == nil {
		t.Fatal("expected an error.")
	}
}

func TestPointSubtraction(t *testing.T) {
	point1 := Point{3, 2, 1}
	point2 := Point{5, 6, 7}
	expected := Vector{-2, -4, -6}

	actual, err := Sub(&point1, &point2)
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Vector{}) {
		t.Fatalf("result is not a Vector. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", &expected, actual)
	}
}

func TestVectorFromPointSubtraction(t *testing.T) {
	point := Point{3, 2, 1}
	vector := Vector{5, 6, 7}
	expected := Point{-2, -4, -6}

	actual, err := Sub(&point, &vector)
	if err != nil {
		t.Fatal("Expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Point{}) {
		t.Fatalf("result is not a Point. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", &expected, actual)
	}
}

func TestVectorSubtraction(t *testing.T) {
	vector1 := Vector{3, 2, 1}
	vector2 := Vector{5, 6, 7}
	expected := Vector{-2, -4, -6}

	actual, err := Sub(&vector1, &vector2)
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Vector{}) {
		t.Fatalf("result is not a Vector. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", &expected, actual)
	}
}

func TestPointFromVectorSubtraction(t *testing.T) {
	vector := Vector{3, 2, 1}
	point := Point{5, 6, 7}

	_, err := Sub(&vector, &point)
	if err == nil {
		t.Fatal("expected an error.")
	}
}

func TestTupleNegationWithSub(t *testing.T) {
	zero := NewZeroVector()
	vector := Vector{1, -2, 3}
	expected := Vector{-1, 2, -3}

	actual, err := Sub(zero, &vector)
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Vector{}) {
		t.Fatalf("result is not a Vector. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", &expected, actual)
	}
}

func TestVectorNegation(t *testing.T) {
	vector := Vector{1, -2, 3}
	expected := Vector{-1, 2, -3}

	actual, err := Negate(&vector)
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Vector{}) {
		t.Fatalf("result is not a Vector. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", &expected, actual)
	}
}

func TestPointNegation(t *testing.T) {
	point := Point{1, -2, 3}
	expected := Point{-1, 2, -3}

	actual, err := Negate(&point)
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Point{}) {
		t.Fatalf("result is not a Point. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", &expected, actual)
	}
}

func TestVectorScalarMultiplication(t *testing.T) {
	vector := Vector{1, -2, 3}
	expected := Vector{3.5, -7, 10.5}

	actual, err := Multiply(&vector, 3.5)
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Vector{}) {
		t.Fatalf("result is not a Vector. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", &expected, actual)
	}
}

func TestPointScalarMultiplication(t *testing.T) {
	point := Point{1, -2, 3}
	expected := Point{3.5, -7, 10.5}

	actual, err := Multiply(&point, 3.5)
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Point{}) {
		t.Fatalf("result is not a Point. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", &expected, actual)
	}
}

func TestVectorFractionMultiplication(t *testing.T) {
	vector := Vector{1, -2, 3}
	expected := Vector{0.5, -1, 1.5}

	actual, err := Multiply(&vector, 0.5)
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Vector{}) {
		t.Fatalf("result is not a Vector. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", &expected, actual)
	}
}

func TestPointFractionMultiplication(t *testing.T) {
	point := Point{1, -2, 3}
	expected := Point{0.5, -1, 1.5}

	actual, err := Multiply(&point, 0.5)
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Point{}) {
		t.Fatalf("result is not a Point. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", &expected, actual)
	}
}

func TestVectorScalarDivision(t *testing.T) {
	vector := Vector{1, -2, 3}
	expected := Vector{0.5, -1, 1.5}

	actual, err := Divide(&vector, 2)
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Vector{}) {
		t.Fatalf("result is not a Vector. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", &expected, actual)
	}
}

func TestPointScalarDivision(t *testing.T) {
	point := Point{1, -2, 3}
	expected := Point{0.5, -1, 1.5}

	actual, err := Divide(&point, 2)
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Point{}) {
		t.Fatalf("result is not a Point. got=%T", actual)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", &expected, actual)
	}
}

func TestVectorMagnitude(t *testing.T) {
	vectors := map[Vector]float64{
		{1, 0, 0}:    1,
		{0, 1, 0}:    1,
		{0, 0, 1}:    1,
		{1, 2, 3}:    math.Sqrt(14),
		{-1, -2, -3}: math.Sqrt(14),
	}

	for v, mag := range vectors {
		actual := Magnitude(&v)

		if !reflect.DeepEqual(actual, mag) {
			t.Fatalf("expected %+v. got=%+v", mag, actual)
		}
	}
}

func TestVectorNormalization(t *testing.T) {
	vectors := map[Vector]Vector{
		{4, 0, 0}: {1, 0, 0},
		{1, 2, 3}: {1 / math.Sqrt(14), 2 / math.Sqrt(14), 3 / math.Sqrt(14)},
	}

	for v, expected := range vectors {
		actual := Normalize(&v)

		if !reflect.DeepEqual(actual, &expected) {
			t.Fatalf("expected %+v. got=%+v", &expected, actual)
		}
	}
}

func TestNormalizedVectorMagnitude(t *testing.T) {
	vector := Vector{1, 2, 3}
	normalized := Normalize(&vector)

	actual := Magnitude(normalized)
	if actual != 1 {
		t.Fatalf("expected 1. got=%+v", actual)
	}
}

func TestTupleDotProduct(t *testing.T) {
	vector1 := Vector{1, 2, 3}
	vector2 := Vector{2, 3, 4}

	actual := Dot(&vector1, &vector2)
	if actual != 20 {
		t.Fatalf("expected 20. got=%f", actual)
	}
}

func TestTupleCrossProduct(t *testing.T) {
	vector1 := Vector{1, 2, 3}
	vector2 := Vector{2, 3, 4}
	expected1 := Vector{-1, 2, -1}
	expected2 := Vector{1, -2, 1}

	actual1 := Cross(&vector1, &vector2)
	if !reflect.DeepEqual(actual1, &expected1) {
		t.Fatalf("expected %+v. got=%+v", &expected1, actual1)
	}

	actual2 := Cross(&vector2, &vector1)
	if !reflect.DeepEqual(actual2, &expected2) {
		t.Fatalf("expected %+v. got=%+v", &expected2, actual2)
	}
}
