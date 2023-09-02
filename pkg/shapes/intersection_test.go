package shapes

import (
	"fmt"
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

type test struct {
	xs     Intersections
	s      Shape
	i1, i2 Intersection
}

func TestIntersectionHit(t *testing.T) {
	s := Sphere()
	i1 := Intersection{T: 1, Object: s}
	i2 := Intersection{T: 2, Object: s}
	xs := Intersections{i2, i1}

	i := xs.Hit()
	if !reflect.DeepEqual(i, i1) {
		t.Fatalf("expected %+v. got=%+v", i1, i)
	}

	s = Sphere()
	i1 = Intersection{T: -1, Object: s}
	i2 = Intersection{T: 1, Object: s}
	xs = Intersections{i2, i1}
	i = xs.Hit()
	if !reflect.DeepEqual(i, i2) {
		t.Fatalf("expected %+v. got=%+v", i2, i)
	}

	s = Sphere()
	i1 = Intersection{T: -2, Object: s}
	i2 = Intersection{T: -1, Object: s}
	xs = Intersections{i2, i1}
	i = xs.Hit()
	fmt.Printf("%+v", i)
	// if i != nil {
	// 	t.Fatalf("expected nil. got=%+v", i)
	// }

	s = Sphere()
	i1 = Intersection{T: 5, Object: s}
	i2 = Intersection{T: 7, Object: s}
	i3 := Intersection{T: -3, Object: s}
	i4 := Intersection{T: 2, Object: s}
	xs = Intersections{i1, i2, i3, i4}
	i = xs.Hit()
	if !reflect.DeepEqual(i, i4) {
		t.Fatalf("expected %+v. got=%+v", i4, i)
	}
}
