package matrix

import (
	"reflect"
	"testing"

	"github.com/adriffaud/ray-tracer-challenge/tuple"
)

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
