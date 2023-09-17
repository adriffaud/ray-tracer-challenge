package internal

import (
	"reflect"
	"testing"
)

func TestIntersectionStruct(t *testing.T) {
	s := Sphere()
	i := Intersection{Distance: 3.5, Object: s}

	if i.Distance != 3.5 {
		t.Fatalf("expected 3.5. got=%f", i.Distance)
	}

	if !reflect.DeepEqual(i.Object, s) {
		t.Fatalf("expected %+v. got=%+v", s, i.Object)
	}
}

func TestIntersectionAggregation(t *testing.T) {
	s := Sphere()
	i1 := Intersection{Distance: 1, Object: s}
	i2 := Intersection{Distance: 2, Object: s}
	xs := Intersections{i1, i2}

	if len(xs) < 2 {
		t.Fatalf("expected 2 values. got=%d", len(xs))
	}

	if xs[0].Distance != 1 {
		t.Fatalf("expected value to be 1. got=%f", xs[0].Distance)
	}
	if xs[1].Distance != 2 {
		t.Fatalf("expected value to be 2. got=%f", xs[0].Distance)
	}
}

type test struct {
	xs     Intersections
	s      Shape
	i1, i2 Intersection
}

func TestIntersectionHit(t *testing.T) {
	s := Sphere()
	i1 := Intersection{Distance: 1, Object: s}
	i2 := Intersection{Distance: 2, Object: s}
	xs := Intersections{i2, i1}

	i, _ := xs.Hit()
	if !reflect.DeepEqual(i, i1) {
		t.Fatalf("expected %+v. got=%+v", i1, i)
	}

	s = Sphere()
	i1 = Intersection{Distance: -1, Object: s}
	i2 = Intersection{Distance: 1, Object: s}
	xs = Intersections{i2, i1}
	i, _ = xs.Hit()
	if !reflect.DeepEqual(i, i2) {
		t.Fatalf("expected %+v. got=%+v", i2, i)
	}

	s = Sphere()
	i1 = Intersection{Distance: -2, Object: s}
	i2 = Intersection{Distance: -1, Object: s}
	xs = Intersections{i2, i1}
	_, hit := xs.Hit()
	if hit {
		t.Fatal("expected no hit")
	}

	s = Sphere()
	i1 = Intersection{Distance: 5, Object: s}
	i2 = Intersection{Distance: 7, Object: s}
	i3 := Intersection{Distance: -3, Object: s}
	i4 := Intersection{Distance: 2, Object: s}
	xs = Intersections{i1, i2, i3, i4}
	i, _ = xs.Hit()
	if !reflect.DeepEqual(i, i4) {
		t.Fatalf("expected %+v. got=%+v", i4, i)
	}
}

func TestIntersectionPrecompute(t *testing.T) {
	r := Ray{
		Origin:    Point{Z: -5},
		Direction: Vector{Z: 1},
	}
	shape := Sphere()
	i := Intersection{Object: shape, Distance: 4}
	comps := i.PrepareComputations(r)

	if comps.Distance != i.Distance {
		t.Fatalf("expected %f. got=%f", i.Distance, comps.Distance)
	}

	if !reflect.DeepEqual(comps.Object, i.Object) {
		t.Fatalf("expected %+v. got=%+v", i.Object, comps.Object)
	}

	p := Point{Z: -1}
	if !reflect.DeepEqual(comps.Point, p) {
		t.Fatalf("expected %+v. got=%+v", p, comps.Point)
	}

	v := Vector{Z: -1}
	if !reflect.DeepEqual(comps.EyeV, v) {
		t.Fatalf("expected %+v. got=%+v", v, comps.EyeV)
	}

	n := Vector{Z: -1}
	if !reflect.DeepEqual(comps.NormalV, n) {
		t.Fatalf("expected %+v. got=%+v", n, comps.NormalV)
	}

	if comps.Inside {
		t.Fatal("expected false")
	}
}

func TestHitInsideIntersection(t *testing.T) {
	r := Ray{
		Origin:    Point{},
		Direction: Vector{Z: 1},
	}
	shape := Sphere()
	i := Intersection{Object: shape, Distance: 1}
	comps := i.PrepareComputations(r)

	p := Point{Z: 1}
	if !reflect.DeepEqual(comps.Point, p) {
		t.Fatalf("expected %+v. got=%+v", p, comps.Point)
	}

	v := Vector{Z: -1}
	if !reflect.DeepEqual(comps.EyeV, v) {
		t.Fatalf("expected %+v. got=%+v", v, comps.EyeV)
	}

	if !comps.Inside {
		t.Fatal("expected to be true")
	}

	n := Vector{Z: -1}
	if !reflect.DeepEqual(comps.NormalV, n) {
		t.Fatalf("expected %+v. got=%+v", n, comps.NormalV)
	}
}
