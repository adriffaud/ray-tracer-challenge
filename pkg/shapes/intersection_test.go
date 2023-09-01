package shapes

import (
	"reflect"
	"testing"
)

func TestIntersectionStruct(t *testing.T) {
	s := Sphere()
	i := Intersection{T: 3.5, Object: s}

	if i.T != 3.5 {
		t.Fatalf("expected 3.5. got=%f", i.T)
	}

	if !reflect.DeepEqual(i.Object, s) {
		t.Fatalf("expected %+v. got=%+v", s, i.Object)
	}
}

func TestIntersectionAggregation(t *testing.T) {
	s := Sphere()
	i1 := Intersection{T: 1, Object: s}
	i2 := Intersection{T: 2, Object: s}
	xs := Intersections{i1, i2}

	if len(xs) < 2 {
		t.Fatalf("expected 2 values. got=%d", len(xs))
	}

	if xs[0].T != 1 {
		t.Fatalf("expected value to be 1. got=%f", xs[0].T)
	}
	if xs[1].T != 2 {
		t.Fatalf("expected value to be 2. got=%f", xs[0].T)
	}
}
