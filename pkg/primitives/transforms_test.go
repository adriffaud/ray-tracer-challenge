package primitives

import (
	"math"
	"reflect"
	"testing"

	"github.com/adriffaud/ray-tracer-challenge/pkg/float"
)

func assertEquals(t *testing.T, expected, actual Tuple) {
	if !float.ApproxEq(expected.X(), actual.X()) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
	if !float.ApproxEq(expected.Y(), actual.Y()) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
	if !float.ApproxEq(expected.Z(), actual.Z()) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestTranslationMultiplication(t *testing.T) {
	transform := Translation(5, -3, 2)
	p := Point{XVal: -3, YVal: 4, ZVal: 5}
	expected := Point{XVal: 2, YVal: 1, ZVal: 7}
	actual, err := transform.MultiplyTuple(&p)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestInverseTranslationMultiplication(t *testing.T) {
	transform := Translation(5, -3, 2)
	inv, err := transform.Inverse()
	if err != nil {
		t.Fatal(err)
	}
	p := Point{XVal: -3, YVal: 4, ZVal: 5}
	expected := Point{XVal: -8, YVal: 7, ZVal: 3}
	actual, err := inv.MultiplyTuple(&p)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestVectorTranslation(t *testing.T) {
	transform := Translation(5, -3, 2)
	v := Vector{XVal: -3, YVal: 4, ZVal: 5}
	actual, err := transform.MultiplyTuple(&v)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, &v) {
		t.Fatalf("expected %+v. got=%+v", &v, actual)
	}
}

func TestPointMatrixScaling(t *testing.T) {
	transform := Scaling(2, 3, 4)
	p := Point{XVal: -4, YVal: 6, ZVal: 8}
	expected := Point{XVal: -8, YVal: 18, ZVal: 32}
	actual, err := transform.MultiplyTuple(&p)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestVectorMatrixScaling(t *testing.T) {
	transform := Scaling(2, 3, 4)
	v := Vector{XVal: -4, YVal: 6, ZVal: 8}
	expected := Vector{XVal: -8, YVal: 18, ZVal: 32}
	actual, err := transform.MultiplyTuple(&v)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestInverseMatrixScaling(t *testing.T) {
	transform := Scaling(2, 3, 4)
	inv, err := transform.Inverse()
	if err != nil {
		t.Fatal(err)
	}
	v := Vector{XVal: -4, YVal: 6, ZVal: 8}
	expected := Vector{XVal: -2, YVal: 2, ZVal: 2}
	actual, err := inv.MultiplyTuple(&v)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestMatrixReflection(t *testing.T) {
	transform := Scaling(-1, 1, 1)
	p := Point{XVal: 2, YVal: 3, ZVal: 4}
	expected := Point{XVal: -2, YVal: 3, ZVal: 4}
	actual, err := transform.MultiplyTuple(&p)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestMatrixRotation(t *testing.T) {
	tests := []struct {
		transform        Matrix
		origin, expected Point
	}{
		{RotationX(math.Pi / 4), Point{XVal: 0, YVal: 1, ZVal: 0}, Point{XVal: 0, YVal: math.Sqrt(2) / 2, ZVal: math.Sqrt(2) / 2}},
		{RotationX(math.Pi / 2), Point{XVal: 0, YVal: 1, ZVal: 0}, Point{XVal: 0, YVal: 0, ZVal: 1}},
		{RotationY(math.Pi / 4), Point{XVal: 0, YVal: 0, ZVal: 1}, Point{XVal: math.Sqrt(2) / 2, YVal: 0, ZVal: math.Sqrt(2) / 2}},
		{RotationY(math.Pi / 2), Point{XVal: 0, YVal: 0, ZVal: 1}, Point{XVal: 1, YVal: 0, ZVal: 0}},
		{RotationZ(math.Pi / 4), Point{XVal: 0, YVal: 1, ZVal: 0}, Point{XVal: -math.Sqrt(2) / 2, YVal: math.Sqrt(2) / 2, ZVal: 0}},
		{RotationZ(math.Pi / 2), Point{XVal: 0, YVal: 1, ZVal: 0}, Point{XVal: -1, YVal: 0, ZVal: 0}},
	}

	for _, test := range tests {
		actual, err := test.transform.MultiplyTuple(&test.origin)
		if err != nil {
			t.Fatal(err)
		}
		assertEquals(t, &test.expected, actual)
	}
}

func TestInverseMatrixXRotation(t *testing.T) {
	p := Point{XVal: 0, YVal: 1, ZVal: 0}
	halfQuarter := RotationX(math.Pi / 4)
	inv, err := halfQuarter.Inverse()
	if err != nil {
		t.Fatal(err)
	}
	expectedHalf := Point{XVal: 0, YVal: math.Sqrt(2) / 2, ZVal: -math.Sqrt(2) / 2}
	actualHalf, err := inv.MultiplyTuple(&p)
	if err != nil {
		t.Fatal(err)
	}
	assertEquals(t, &expectedHalf, actualHalf)
}

func TestShearing(t *testing.T) {
	tests := []struct {
		transform        Matrix
		origin, expected Point
	}{
		{Shearing(1, 0, 0, 0, 0, 0), Point{XVal: 2, YVal: 3, ZVal: 4}, Point{XVal: 5, YVal: 3, ZVal: 4}},
		{Shearing(0, 1, 0, 0, 0, 0), Point{XVal: 2, YVal: 3, ZVal: 4}, Point{XVal: 6, YVal: 3, ZVal: 4}},
		{Shearing(0, 0, 1, 0, 0, 0), Point{XVal: 2, YVal: 3, ZVal: 4}, Point{XVal: 2, YVal: 5, ZVal: 4}},
		{Shearing(0, 0, 0, 1, 0, 0), Point{XVal: 2, YVal: 3, ZVal: 4}, Point{XVal: 2, YVal: 7, ZVal: 4}},
		{Shearing(0, 0, 0, 0, 1, 0), Point{XVal: 2, YVal: 3, ZVal: 4}, Point{XVal: 2, YVal: 3, ZVal: 6}},
		{Shearing(0, 0, 0, 0, 0, 1), Point{XVal: 2, YVal: 3, ZVal: 4}, Point{XVal: 2, YVal: 3, ZVal: 7}},
	}

	for _, test := range tests {
		actual, err := test.transform.MultiplyTuple(&test.origin)
		if err != nil {
			t.Fatal(err)
		}

		assertEquals(t, &test.expected, actual)
	}
}

func TestTransformationSequence(t *testing.T) {
	p := Point{XVal: 1, YVal: 0, ZVal: 1}
	a := RotationX(math.Pi / 2)
	b := Scaling(5, 5, 5)
	c := Translation(10, 5, 7)

	expected := Point{XVal: 1, YVal: -1, ZVal: 0}
	p2, err := a.MultiplyTuple(&p)
	if err != nil {
		t.Fatal(err)
	}
	assertEquals(t, &expected, p2)

	expected = Point{XVal: 5, YVal: -5, ZVal: 0}
	p3, err := b.MultiplyTuple(p2)
	if err != nil {
		t.Fatal(err)
	}
	assertEquals(t, &expected, p3)

	expected = Point{XVal: 15, YVal: 0, ZVal: 7}
	p4, err := c.MultiplyTuple(p3)
	if err != nil {
		t.Fatal(err)
	}
	assertEquals(t, &expected, p4)
}

func TestChainedTransformations(t *testing.T) {
	p := Point{XVal: 1, YVal: 0, ZVal: 1}
	a := RotationX(math.Pi / 2)
	b := Scaling(5, 5, 5)
	c := Translation(10, 5, 7)

	transform := c.Multiply(b.Multiply(a))
	expected := Point{XVal: 15, YVal: 0, ZVal: 7}
	actual, err := transform.MultiplyTuple(&p)
	if err != nil {
		t.Fatal(err)
	}
	assertEquals(t, &expected, actual)
}
