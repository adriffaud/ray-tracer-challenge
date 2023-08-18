package matrix

import (
	"reflect"
	"testing"

	"github.com/adriffaud/ray-tracer-challenge/tuple"
)

func Test4x4MatrixConstruction(t *testing.T) {
	m := Matrix{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	}

	checks := map[[2]int]float64{
		{0, 0}: 1,
		{0, 3}: 4,
		{1, 0}: 5.5,
		{1, 2}: 7.5,
		{2, 2}: 11,
		{3, 0}: 13.5,
		{3, 2}: 15.5,
	}

	for coords, expected := range checks {
		actual := m[coords[0]][coords[1]]
		if actual != expected {
			t.Fatalf("expected m[%d,%d] to equal %f. got=%f", coords[0], coords[1], expected, actual)
		}
	}
}

func Test2x2MatrixCreation(t *testing.T) {
	m := Matrix{
		{-3, 5},
		{1, -2},
	}

	checks := map[[2]int]float64{
		{0, 0}: -3,
		{0, 1}: 5,
		{1, 0}: 1,
		{1, 1}: -2,
	}

	for coords, expected := range checks {
		actual := m[coords[0]][coords[1]]
		if actual != expected {
			t.Fatalf("expected m[%d,%d] to equal %f. got=%f", coords[0], coords[1], expected, actual)
		}
	}
}

func Test3x3MatrixCreation(t *testing.T) {
	m := Matrix{
		{-3, 5, 0},
		{1, -2, -7},
		{0, 1, 1},
	}

	checks := map[[2]int]float64{
		{0, 0}: -3,
		{1, 1}: -2,
		{2, 2}: 1,
	}

	for coords, expected := range checks {
		actual := m[coords[0]][coords[1]]
		if actual != expected {
			t.Fatalf("expected m[%d,%d] to equal %f. got=%f", coords[0], coords[1], expected, actual)
		}
	}
}

func TestMatrixEquality(t *testing.T) {
	a := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	b := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}

	if !reflect.DeepEqual(a, b) {
		t.Fatalf("expected matrices to be equal. got=\n%+v\n%+v", a, b)
	}
}

func TestMatrixInequality(t *testing.T) {
	a := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	b := Matrix{
		{2, 3, 4, 5},
		{6, 7, 8, 9},
		{8, 7, 6, 5},
		{4, 3, 2, 1},
	}

	if reflect.DeepEqual(a, b) {
		t.Fatalf("expected matrices to be different. got=\n%+v\n%+v", a, b)
	}
}

func TestMatrixMultiplication(t *testing.T) {
	a := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	b := Matrix{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	}
	expected := Matrix{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	}

	actual := Multiply(a, b)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected matrices to be equal. got=\n%+v", actual)
	}
}

func TestMatrixTupleMultiplication(t *testing.T) {
	a := Matrix{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	}
	expected := tuple.Point{XVal: 18, YVal: 24, ZVal: 33}

	actual, err := MultiplyTuple(a, &tuple.Point{XVal: 1, YVal: 2, ZVal: 3})
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&tuple.Point{}) {
		t.Fatalf("expected %T. got=%T", reflect.TypeOf(&tuple.Point{}), reflect.TypeOf(actual))
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", &expected, actual)
	}
}
