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

func TestMatrixIdentityMultiplication(t *testing.T) {
	a := Matrix{
		{0, 1, 2, 4},
		{1, 2, 4, 8},
		{2, 4, 8, 16},
		{4, 8, 16, 32},
	}

	actual := Multiply(a, IdentityMatrix())
	if !reflect.DeepEqual(actual, a) {
		t.Fatalf("expected %+v. got=%+v", a, actual)
	}
}

func TestMatrixTransposition(t *testing.T) {
	a := Matrix{
		{0, 9, 3, 0},
		{9, 8, 0, 8},
		{1, 8, 5, 3},
		{0, 0, 5, 8},
	}
	expected := Matrix{
		{0, 9, 1, 0},
		{9, 8, 8, 0},
		{3, 0, 5, 5},
		{0, 8, 3, 8},
	}

	actual := Transpose(a)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestIdentityMatrixTransposition(t *testing.T) {
	a := IdentityMatrix()
	actual := Transpose(a)
	if !reflect.DeepEqual(actual, a) {
		t.Fatalf("expected %+v. got=%+v", a, actual)
	}
}

func Test2x2Determinant(t *testing.T) {
	a := Matrix{
		{1, 5},
		{-3, 2},
	}

	actual := Determinant(a)
	if actual != 17 {
		t.Fatalf("expected 17. got=%d", actual)
	}
}

func Test3x3Submatrix(t *testing.T) {
	a := Matrix{
		{1, 5, 0},
		{-3, 2, 7},
		{0, 6, -3},
	}
	expected := Matrix{
		{-3, 2},
		{0, 6},
	}

	actual := Submatrix(a, 0, 2)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func Test4x4Submatrix(t *testing.T) {
	a := Matrix{
		{-6, 1, 1, 6},
		{-8, 5, 8, 6},
		{-1, 0, 8, 2},
		{-7, 1, -1, 1},
	}
	expected := Matrix{
		{-6, 1, 6},
		{-8, 8, 6},
		{-7, -1, 1},
	}

	actual := Submatrix(a, 2, 1)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func Test3x3Minor(t *testing.T) {
	a := Matrix{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	}
	b := Submatrix(a, 1, 0)
	determinant := Determinant(b)
	if determinant != 25 {
		t.Fatalf("expected determinant to be 25. got=%d", determinant)
	}

	minor := Minor(a, 1, 0)
	if minor != 25 {
		t.Fatalf("expected minor to be 25. got=%d", minor)
	}
}

func Test3x3Cofactor(t *testing.T) {
	a := Matrix{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	}

	minor := Minor(a, 0, 0)
	if minor != -12 {
		t.Fatalf("expected minor=-12. got=%d", minor)
	}

	cofactor := Cofactor(a, 0, 0)
	if cofactor != -12 {
		t.Fatalf("expected cofactor=-12. got=%d", cofactor)
	}

	minor = Minor(a, 1, 0)
	if minor != 25 {
		t.Fatalf("expected minor=25. got=%d", minor)
	}

	cofactor = Cofactor(a, 1, 0)
	if cofactor != -25 {
		t.Fatalf("expected cofactor=-25. got=%d", cofactor)
	}
}

func Test3x3Determinant(t *testing.T) {
	a := Matrix{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	}
	cofactor := Cofactor(a, 0, 0)
	if cofactor != 56 {
		t.Fatalf("expected 56.got=%d", cofactor)
	}
	cofactor = Cofactor(a, 0, 1)
	if cofactor != 12 {
		t.Fatalf("expected 12.got=%d", cofactor)
	}
	cofactor = Cofactor(a, 0, 2)
	if cofactor != -46 {
		t.Fatalf("expected -46.got=%d", cofactor)
	}
	determinant := Determinant(a)
	if determinant != -196 {
		t.Fatalf("expected -196. got=%d", determinant)
	}
}

func Test4x4Determinant(t *testing.T) {
	a := Matrix{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	}
	cofactor := Cofactor(a, 0, 0)
	if cofactor != 690 {
		t.Fatalf("expected 690.got=%d", cofactor)
	}
	cofactor = Cofactor(a, 0, 1)
	if cofactor != 447 {
		t.Fatalf("expected 447.got=%d", cofactor)
	}
	cofactor = Cofactor(a, 0, 2)
	if cofactor != 210 {
		t.Fatalf("expected 210.got=%d", cofactor)
	}
	cofactor = Cofactor(a, 0, 3)
	if cofactor != 51 {
		t.Fatalf("expected 51.got=%d", cofactor)
	}
	determinant := Determinant(a)
	if determinant != -4071 {
		t.Fatalf("expected -4071. got=%d", determinant)
	}
}
