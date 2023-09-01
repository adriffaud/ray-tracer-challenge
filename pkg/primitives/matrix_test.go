package primitives

import (
	"reflect"
	"testing"
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

	actual := a.Multiply(b)

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
	expected := Point{XVal: 18, YVal: 24, ZVal: 33}

	actual, err := a.MultiplyTuple(&Point{XVal: 1, YVal: 2, ZVal: 3})
	if err != nil {
		t.Fatal("expected no error.")
	}

	if reflect.TypeOf(actual) != reflect.TypeOf(&Point{}) {
		t.Fatalf("expected %T. got=%T", reflect.TypeOf(&Point{}), reflect.TypeOf(actual))
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

	actual := IdentityMatrix().Multiply(a)
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

	actual := a.Transpose()
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestIdentityMatrixTransposition(t *testing.T) {
	a := IdentityMatrix()
	actual := a.Transpose()
	if !reflect.DeepEqual(actual, a) {
		t.Fatalf("expected %+v. got=%+v", a, actual)
	}
}

func Test2x2Determinant(t *testing.T) {
	a := Matrix{
		{1, 5},
		{-3, 2},
	}

	actual := a.Determinant()
	if actual != 17 {
		t.Fatalf("expected 17. got=%f", actual)
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

	actual := a.Submatrix(0, 2)
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

	actual := a.Submatrix(2, 1)
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
	b := a.Submatrix(1, 0)
	determinant := b.Determinant()
	if determinant != 25 {
		t.Fatalf("expected determinant to be 25. got=%f", determinant)
	}

	minor := a.Minor(1, 0)
	if minor != 25 {
		t.Fatalf("expected minor to be 25. got=%f", minor)
	}
}

func Test3x3Cofactor(t *testing.T) {
	a := Matrix{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	}

	minor := a.Minor(0, 0)
	if minor != -12 {
		t.Fatalf("expected minor=-12. got=%f", minor)
	}

	cofactor := a.Cofactor(0, 0)
	if cofactor != -12 {
		t.Fatalf("expected cofactor=-12. got=%f", cofactor)
	}

	minor = a.Minor(1, 0)
	if minor != 25 {
		t.Fatalf("expected minor=25. got=%f", minor)
	}

	cofactor = a.Cofactor(1, 0)
	if cofactor != -25 {
		t.Fatalf("expected cofactor=-25. got=%f", cofactor)
	}
}

func Test3x3Determinant(t *testing.T) {
	a := Matrix{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	}
	cofactor := a.Cofactor(0, 0)
	if cofactor != 56 {
		t.Fatalf("expected 56.got=%f", cofactor)
	}
	cofactor = a.Cofactor(0, 1)
	if cofactor != 12 {
		t.Fatalf("expected 12.got=%f", cofactor)
	}
	cofactor = a.Cofactor(0, 2)
	if cofactor != -46 {
		t.Fatalf("expected -46.got=%f", cofactor)
	}
	determinant := a.Determinant()
	if determinant != -196 {
		t.Fatalf("expected -196. got=%f", determinant)
	}
}

func Test4x4Determinant(t *testing.T) {
	a := Matrix{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	}
	cofactor := a.Cofactor(0, 0)
	if cofactor != 690 {
		t.Fatalf("expected 690.got=%f", cofactor)
	}
	cofactor = a.Cofactor(0, 1)
	if cofactor != 447 {
		t.Fatalf("expected 447.got=%f", cofactor)
	}
	cofactor = a.Cofactor(0, 2)
	if cofactor != 210 {
		t.Fatalf("expected 210.got=%f", cofactor)
	}
	cofactor = a.Cofactor(0, 3)
	if cofactor != 51 {
		t.Fatalf("expected 51.got=%f", cofactor)
	}
	determinant := a.Determinant()
	if determinant != -4071 {
		t.Fatalf("expected -4071. got=%f", determinant)
	}
}

func TestInvertibleMatrix(t *testing.T) {
	a := Matrix{
		{6, 4, 4, 4},
		{5, 5, 7, 6},
		{4, -9, 3, -7},
		{9, 1, 7, -6},
	}

	det := a.Determinant()
	if det != -2120 {
		t.Fatalf("expected -2120. got=%f", det)
	}
}

func TestNoninvertibleMatrix(t *testing.T) {
	a := Matrix{
		{-4, 2, -2, -3},
		{9, 6, 2, 6},
		{0, -5, 1, -5},
		{0, 0, 0, 0},
	}

	det := a.Determinant()
	if det != 0 {
		t.Fatalf("expected 0. got=%f", det)
	}
}

func TestMatrixInverse(t *testing.T) {
	a := Matrix{
		{-5, 2, 6, -8},
		{1, -5, 1, 8},
		{7, 7, -6, -7},
		{1, -3, 7, 4},
	}
	b, err := a.Inverse()
	if err != nil {
		t.Fatal("expected an invertible matrix")
	}

	det := a.Determinant()
	if det != 532 {
		t.Fatalf("expected 532. got=%f", det)
	}

	cofactor := a.Cofactor(2, 3)
	if cofactor != -160 {
		t.Fatalf("expected -160. got=%f", cofactor)
	}

	if b[3][2] != -160.0/532.0 {
		t.Fatal("expected b[3][2] to equal -160/532")
	}

	cofactor = a.Cofactor(3, 2)
	if cofactor != 105 {
		t.Fatalf("expected 105. got=%f", cofactor)
	}

	if b[2][3] != 105.0/532.0 {
		t.Fatal("expected b[2][3] to equal 105/532")
	}

	expected := Matrix{
		{0.21805, 0.45113, 0.24060, -0.04511},
		{-0.80827, -1.45677, -0.44361, 0.52068},
		{-0.07895, -0.22368, -0.05263, 0.19737},
		{-0.52256, -0.81391, -0.30075, 0.30639},
	}

	if !Eq(expected, b) {
		t.Fatalf("expected %+v. got=%+v", expected, b)
	}
}

func TestMatrixInverse2(t *testing.T) {
	a := Matrix{
		{8, -5, 9, 2},
		{7, 5, 6, 1},
		{-6, 0, 9, 6},
		{-3, 0, -9, -4},
	}
	expected := Matrix{
		{-0.15385, -0.15385, -0.28205, -0.53846},
		{-0.07692, 0.12308, 0.02564, 0.03077},
		{0.35897, 0.35897, 0.43590, 0.92308},
		{-0.69231, -0.69231, -0.76923, -1.92308},
	}

	actual, err := a.Inverse()
	if err != nil {
		t.Fatal("expected an invertible matrix")
	}

	if !Eq(expected, actual) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestMatrixInverse3(t *testing.T) {
	a := Matrix{
		{9, 3, 0, 9},
		{-5, -2, -6, -3},
		{-4, 9, 6, 4},
		{-7, 6, 6, 2},
	}
	expected := Matrix{
		{-0.04074, -0.07778, 0.14444, -0.22222},
		{-0.07778, 0.03333, 0.36667, -0.33333},
		{-0.02901, -0.14630, -0.10926, 0.12963},
		{0.17778, 0.06667, -0.26667, 0.33333},
	}

	actual, err := a.Inverse()
	if err != nil {
		t.Fatal(err)
	}

	if !Eq(expected, actual) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestMatrixInverseMultiplication(t *testing.T) {
	a := Matrix{
		{3, -9, 7, 3},
		{3, -8, 2, -9},
		{-4, 4, 4, 1},
		{-6, 5, -1, 1},
	}
	b := Matrix{
		{8, 2, 2, 2},
		{3, -1, 7, 0},
		{7, 0, 5, 4},
		{6, -2, 0, 5},
	}
	c := a.Multiply(b)

	inv, err := b.Inverse()
	if err != nil {
		t.Fatal(err)
	}
	actual := c.Multiply(inv)

	if !Eq(actual, a) {
		t.Fatalf("expected %+v. got=%+v", a, actual)
	}
}
