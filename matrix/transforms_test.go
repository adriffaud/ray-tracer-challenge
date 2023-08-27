package matrix

import (
	"math"
	"reflect"
	"testing"

	"github.com/adriffaud/ray-tracer-challenge/float"
	"github.com/adriffaud/ray-tracer-challenge/tuple"
)

func assertEquals(t *testing.T, expected, actual tuple.Tuple) {
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
	p := tuple.Point{XVal: -3, YVal: 4, ZVal: 5}
	expected := tuple.Point{XVal: 2, YVal: 1, ZVal: 7}
	actual, err := MultiplyTuple(transform, &p)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestInverseTranslationMultiplication(t *testing.T) {
	transform := Translation(5, -3, 2)
	inv, err := Inverse(transform)
	if err != nil {
		t.Fatal(err)
	}
	p := tuple.Point{XVal: -3, YVal: 4, ZVal: 5}
	expected := tuple.Point{XVal: -8, YVal: 7, ZVal: 3}
	actual, err := MultiplyTuple(inv, &p)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestVectorTranslation(t *testing.T) {
	transform := Translation(5, -3, 2)
	v := tuple.Vector{XVal: -3, YVal: 4, ZVal: 5}
	actual, err := MultiplyTuple(transform, &v)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, &v) {
		t.Fatalf("expected %+v. got=%+v", &v, actual)
	}
}

func TestPointMatrixScaling(t *testing.T) {
	transform := Scaling(2, 3, 4)
	p := tuple.Point{XVal: -4, YVal: 6, ZVal: 8}
	expected := tuple.Point{XVal: -8, YVal: 18, ZVal: 32}
	actual, err := MultiplyTuple(transform, &p)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestVectorMatrixScaling(t *testing.T) {
	transform := Scaling(2, 3, 4)
	v := tuple.Vector{XVal: -4, YVal: 6, ZVal: 8}
	expected := tuple.Vector{XVal: -8, YVal: 18, ZVal: 32}
	actual, err := MultiplyTuple(transform, &v)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestInverseMatrixScaling(t *testing.T) {
	transform := Scaling(2, 3, 4)
	inv, err := Inverse(transform)
	if err != nil {
		t.Fatal(err)
	}
	v := tuple.Vector{XVal: -4, YVal: 6, ZVal: 8}
	expected := tuple.Vector{XVal: -2, YVal: 2, ZVal: 2}
	actual, err := MultiplyTuple(inv, &v)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestMatrixReflection(t *testing.T) {
	transform := Scaling(-1, 1, 1)
	p := tuple.Point{XVal: 2, YVal: 3, ZVal: 4}
	expected := tuple.Point{XVal: -2, YVal: 3, ZVal: 4}
	actual, err := MultiplyTuple(transform, &p)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, &expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestMatrixXRotation(t *testing.T) {
	p := tuple.Point{XVal: 0, YVal: 1, ZVal: 0}
	halfQuarter := RotationX(math.Pi / 4)
	expectedHalf := tuple.Point{XVal: 0, YVal: math.Sqrt(2) / 2, ZVal: math.Sqrt(2) / 2}
	actualHalf, err := MultiplyTuple(halfQuarter, &p)
	if err != nil {
		t.Fatal(err)
	}
	assertEquals(t, &expectedHalf, actualHalf)

	fullQuarter := RotationX(math.Pi / 2)
	expectedFull := tuple.Point{XVal: 0, YVal: 0, ZVal: 1}
	actualFull, err := MultiplyTuple(fullQuarter, &p)
	if err != nil {
		t.Fatal(err)
	}
	assertEquals(t, &expectedFull, actualFull)
}

func TestInverseMatrixXRotation(t *testing.T) {
	p := tuple.Point{XVal: 0, YVal: 1, ZVal: 0}
	halfQuarter := RotationX(math.Pi / 4)
	inv, err := Inverse(halfQuarter)
	if err != nil {
		t.Fatal(err)
	}
	expectedHalf := tuple.Point{XVal: 0, YVal: math.Sqrt(2) / 2, ZVal: -math.Sqrt(2) / 2}
	actualHalf, err := MultiplyTuple(inv, &p)
	if err != nil {
		t.Fatal(err)
	}
	assertEquals(t, &expectedHalf, actualHalf)
}

func TestMatrixYRotation(t *testing.T) {
	p := tuple.Point{XVal: 0, YVal: 0, ZVal: 1}
	halfQuarter := RotationY(math.Pi / 4)
	expectedHalf := tuple.Point{XVal: math.Sqrt(2) / 2, YVal: 0, ZVal: math.Sqrt(2) / 2}
	actualHalf, err := MultiplyTuple(halfQuarter, &p)
	if err != nil {
		t.Fatal(err)
	}
	assertEquals(t, &expectedHalf, actualHalf)

	fullQuarter := RotationY(math.Pi / 2)
	expectedFull := tuple.Point{XVal: 1, YVal: 0, ZVal: 0}
	actualFull, err := MultiplyTuple(fullQuarter, &p)
	if err != nil {
		t.Fatal(err)
	}
	assertEquals(t, &expectedFull, actualFull)
}

func TestMatrixZRotation(t *testing.T) {
	p := tuple.Point{XVal: 0, YVal: 1, ZVal: 0}
	halfQuarter := RotationZ(math.Pi / 4)
	expectedHalf := tuple.Point{XVal: -math.Sqrt(2) / 2, YVal: math.Sqrt(2) / 2, ZVal: 0}
	actualHalf, err := MultiplyTuple(halfQuarter, &p)
	if err != nil {
		t.Fatal(err)
	}
	assertEquals(t, &expectedHalf, actualHalf)

	fullQuarter := RotationZ(math.Pi / 2)
	expectedFull := tuple.Point{XVal: -1, YVal: 0, ZVal: 0}
	actualFull, err := MultiplyTuple(fullQuarter, &p)
	if err != nil {
		t.Fatal(err)
	}
	assertEquals(t, &expectedFull, actualFull)
}
