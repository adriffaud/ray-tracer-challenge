package internal

import (
	"math"
	"reflect"
	"testing"
)

func TestTupleAddition(t *testing.T) {
	point := Point{3, -2, 5}
	vector := Vector{-2, 3, 1}
	expected := Point{1, 1, 6}
	actual := point.Add(vector)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestVectorAddition(t *testing.T) {
	vector1 := Vector{3, -2, 5}
	vector2 := Vector{-2, 3, 1}
	expected := Vector{1, 1, 6}
	actual := vector1.Add(vector2)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestPointSubtraction(t *testing.T) {
	point1 := Point{3, 2, 1}
	point2 := Point{5, 6, 7}
	expected := Vector{-2, -4, -6}
	actual := point1.SubPoint(point2)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestVectorFromPointSubtraction(t *testing.T) {
	point := Point{3, 2, 1}
	vector := Vector{5, 6, 7}
	expected := Point{-2, -4, -6}
	actual := point.SubVector(vector)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestVectorSubtraction(t *testing.T) {
	vector1 := Vector{3, 2, 1}
	vector2 := Vector{5, 6, 7}
	expected := Vector{-2, -4, -6}
	actual := vector1.Sub(vector2)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestTupleNegationWithSub(t *testing.T) {
	zero := Vector{}
	vector := Vector{1, -2, 3}
	expected := Vector{-1, 2, -3}
	actual := zero.Sub(vector)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestVectorNegation(t *testing.T) {
	vector := Vector{1, -2, 3}
	expected := Vector{-1, 2, -3}
	actual := vector.Negate()

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestPointNegation(t *testing.T) {
	point := Point{1, -2, 3}
	expected := Point{-1, 2, -3}
	actual := point.Negate()

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestVectorScalarMultiplication(t *testing.T) {
	vector := Vector{1, -2, 3}
	expected := Vector{3.5, -7, 10.5}
	actual := vector.Multiply(3.5)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestPointScalarMultiplication(t *testing.T) {
	point := Point{1, -2, 3}
	expected := Point{3.5, -7, 10.5}
	actual := point.Multiply(3.5)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestVectorFractionMultiplication(t *testing.T) {
	vector := Vector{1, -2, 3}
	expected := Vector{0.5, -1, 1.5}
	actual := vector.Multiply(0.5)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestPointFractionMultiplication(t *testing.T) {
	point := Point{1, -2, 3}
	expected := Point{0.5, -1, 1.5}
	actual := point.Multiply(0.5)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestVectorScalarDivision(t *testing.T) {
	vector := Vector{1, -2, 3}
	expected := Vector{0.5, -1, 1.5}
	actual := vector.Divide(2)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
	}
}

func TestPointScalarDivision(t *testing.T) {
	point := Point{1, -2, 3}
	expected := Point{0.5, -1, 1.5}
	actual := point.Divide(2)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v. got=%+v", expected, actual)
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
		actual := v.Magnitude()

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
		actual := v.Normalize()

		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("expected %+v. got=%+v", expected, actual)
		}
	}
}

func TestNormalizedVectorMagnitude(t *testing.T) {
	vector := Vector{1, 2, 3}
	normalized := vector.Normalize()
	actual := normalized.Magnitude()
	if actual != 1 {
		t.Fatalf("expected 1. got=%+v", actual)
	}
}

func TestTupleDotProduct(t *testing.T) {
	vector1 := Vector{1, 2, 3}
	vector2 := Vector{2, 3, 4}

	actual := vector1.Dot(vector2)
	if actual != 20 {
		t.Fatalf("expected 20. got=%f", actual)
	}
}

func TestTupleCrossProduct(t *testing.T) {
	vector1 := Vector{1, 2, 3}
	vector2 := Vector{2, 3, 4}
	expected1 := Vector{-1, 2, -1}
	expected2 := Vector{1, -2, 1}

	actual1 := vector1.Cross(vector2)
	if !reflect.DeepEqual(actual1, expected1) {
		t.Fatalf("expected %+v. got=%+v", expected1, actual1)
	}

	actual2 := vector2.Cross(vector1)
	if !reflect.DeepEqual(actual2, expected2) {
		t.Fatalf("expected %+v. got=%+v", expected2, actual2)
	}
}

func TestVectorReflection(t *testing.T) {
	tests := []struct {
		in, n, expected Vector
	}{
		{Vector{X: 1, Y: -1}, Vector{Y: 1}, Vector{X: 1, Y: 1}},
		{Vector{Y: -1}, Vector{X: math.Sqrt(2) / 2, Y: math.Sqrt(2) / 2}, Vector{X: 1}},
	}

	for _, test := range tests {
		r := test.in.Reflect(test.n)
		assertVectorEquals(t, test.expected, r)
	}
}
