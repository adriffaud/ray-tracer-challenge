package internal

import (
	"math"
	"reflect"
	"testing"
)

func assertPointEquals(t *testing.T, expected, actual Point) {
	if !ApproxEq(expected.X, actual.X) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
	if !ApproxEq(expected.Y, actual.Y) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
	if !ApproxEq(expected.Z, actual.Z) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestTranslationMultiplication(t *testing.T) {
	transform := Translation(5, -3, 2)
	p := Point{X: -3, Y: 4, Z: 5}
	expected := Point{X: 2, Y: 1, Z: 7}
	actual := p.MultiplyMatrix(transform)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestInverseTranslationMultiplication(t *testing.T) {
	transform := Translation(5, -3, 2)
	inv, err := transform.Inverse()
	if err != nil {
		t.Fatal(err)
	}
	p := Point{X: -3, Y: 4, Z: 5}
	expected := Point{X: -8, Y: 7, Z: 3}
	actual := p.MultiplyMatrix(inv)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestVectorTranslation(t *testing.T) {
	transform := Translation(5, -3, 2)
	v := Vector{X: -3, Y: 4, Z: 5}
	actual := v.MultiplyMatrix(transform)

	if !reflect.DeepEqual(actual, v) {
		t.Fatalf("expected %+v. got=%+v", v, actual)
	}
}

func TestPointMatrixScaling(t *testing.T) {
	transform := Scaling(2, 3, 4)
	p := Point{X: -4, Y: 6, Z: 8}
	expected := Point{X: -8, Y: 18, Z: 32}
	actual := p.MultiplyMatrix(transform)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestVectorMatrixScaling(t *testing.T) {
	transform := Scaling(2, 3, 4)
	v := Vector{X: -4, Y: 6, Z: 8}
	expected := Vector{X: -8, Y: 18, Z: 32}
	actual := v.MultiplyMatrix(transform)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestInverseMatrixScaling(t *testing.T) {
	transform := Scaling(2, 3, 4)
	inv, err := transform.Inverse()
	if err != nil {
		t.Fatal(err)
	}
	v := Vector{X: -4, Y: 6, Z: 8}
	expected := Vector{X: -2, Y: 2, Z: 2}
	actual := v.MultiplyMatrix(inv)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestMatrixReflection(t *testing.T) {
	transform := Scaling(-1, 1, 1)
	p := Point{X: 2, Y: 3, Z: 4}
	expected := Point{X: -2, Y: 3, Z: 4}
	actual := p.MultiplyMatrix(transform)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestMatrixRotation(t *testing.T) {
	tests := []struct {
		transform        Matrix
		origin, expected Point
	}{
		{RotationX(math.Pi / 4), Point{X: 0, Y: 1, Z: 0}, Point{X: 0, Y: math.Sqrt(2) / 2, Z: math.Sqrt(2) / 2}},
		{RotationX(math.Pi / 2), Point{X: 0, Y: 1, Z: 0}, Point{X: 0, Y: 0, Z: 1}},
		{RotationY(math.Pi / 4), Point{X: 0, Y: 0, Z: 1}, Point{X: math.Sqrt(2) / 2, Y: 0, Z: math.Sqrt(2) / 2}},
		{RotationY(math.Pi / 2), Point{X: 0, Y: 0, Z: 1}, Point{X: 1, Y: 0, Z: 0}},
		{RotationZ(math.Pi / 4), Point{X: 0, Y: 1, Z: 0}, Point{X: -math.Sqrt(2) / 2, Y: math.Sqrt(2) / 2, Z: 0}},
		{RotationZ(math.Pi / 2), Point{X: 0, Y: 1, Z: 0}, Point{X: -1, Y: 0, Z: 0}},
	}

	for _, test := range tests {
		actual := test.origin.MultiplyMatrix(test.transform)
		assertPointEquals(t, test.expected, actual)
	}
}

func TestInverseMatrixXRotation(t *testing.T) {
	p := Point{X: 0, Y: 1, Z: 0}
	halfQuarter := RotationX(math.Pi / 4)
	inv, err := halfQuarter.Inverse()
	if err != nil {
		t.Fatal(err)
	}
	expectedHalf := Point{X: 0, Y: math.Sqrt(2) / 2, Z: -math.Sqrt(2) / 2}
	actualHalf := p.MultiplyMatrix(inv)
	assertPointEquals(t, expectedHalf, actualHalf)
}

func TestShearing(t *testing.T) {
	tests := []struct {
		transform        Matrix
		origin, expected Point
	}{
		{Shearing(1, 0, 0, 0, 0, 0), Point{X: 2, Y: 3, Z: 4}, Point{X: 5, Y: 3, Z: 4}},
		{Shearing(0, 1, 0, 0, 0, 0), Point{X: 2, Y: 3, Z: 4}, Point{X: 6, Y: 3, Z: 4}},
		{Shearing(0, 0, 1, 0, 0, 0), Point{X: 2, Y: 3, Z: 4}, Point{X: 2, Y: 5, Z: 4}},
		{Shearing(0, 0, 0, 1, 0, 0), Point{X: 2, Y: 3, Z: 4}, Point{X: 2, Y: 7, Z: 4}},
		{Shearing(0, 0, 0, 0, 1, 0), Point{X: 2, Y: 3, Z: 4}, Point{X: 2, Y: 3, Z: 6}},
		{Shearing(0, 0, 0, 0, 0, 1), Point{X: 2, Y: 3, Z: 4}, Point{X: 2, Y: 3, Z: 7}},
	}

	for _, test := range tests {
		actual := test.origin.MultiplyMatrix(test.transform)
		assertPointEquals(t, test.expected, actual)
	}
}

func TestTransformationSequence(t *testing.T) {
	p := Point{X: 1, Y: 0, Z: 1}
	a := RotationX(math.Pi / 2)
	b := Scaling(5, 5, 5)
	c := Translation(10, 5, 7)

	expected := Point{X: 1, Y: -1, Z: 0}
	p2 := p.MultiplyMatrix(a)
	assertPointEquals(t, expected, p2)

	expected = Point{X: 5, Y: -5, Z: 0}
	p3 := p2.MultiplyMatrix(b)
	assertPointEquals(t, expected, p3)

	expected = Point{X: 15, Y: 0, Z: 7}
	p4 := p3.MultiplyMatrix(c)
	assertPointEquals(t, expected, p4)
}

func TestChainedTransformations(t *testing.T) {
	p := Point{X: 1, Y: 0, Z: 1}
	a := RotationX(math.Pi / 2)
	b := Scaling(5, 5, 5)
	c := Translation(10, 5, 7)

	transform := c.Multiply(b.Multiply(a))
	expected := Point{X: 15, Y: 0, Z: 7}
	actual := p.MultiplyMatrix(transform)
	assertPointEquals(t, expected, actual)
}

func TestViewTransformationMatrix(t *testing.T) {
	tests := []struct {
		expected Matrix
		from, to Point
		up       Vector
	}{
		{IdentityMatrix(), Point{}, Point{Z: -1}, Vector{Y: 1}},
		{Scaling(-1, 1, -1), Point{}, Point{Z: 1}, Vector{Y: 1}},
		{Translation(0, 0, -8), Point{Z: 8}, Point{}, Vector{Y: 1}},
		{
			Matrix{
				{-0.50709, 0.50709, 0.67612, -2.36643},
				{0.76772, 0.60609, 0.12122, -2.82843},
				{-0.35857, 0.59761, -0.71714, 0},
				{0, 0, 0, 1},
			},
			Point{X: 1, Y: 3, Z: 2},
			Point{X: 4, Y: -2, Z: 8},
			Vector{X: 1, Y: 1},
		},
	}

	for _, test := range tests {
		transform := ViewTransform(test.from, test.to, test.up)

		if !Eq(transform, test.expected) {
			t.Fatalf("expected %+v. got=%+v", test.expected, transform)
		}

	}
}
