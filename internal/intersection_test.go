package internal

import (
	"reflect"
	"testing"
)

func TestIntersectionStruct(t *testing.T) {
	s := Sphere()
	i := Intersection{s, 3.5}

	if i.Distance != 3.5 {
		t.Fatalf("expected 3.5. got=%f", i.Distance)
	}

	if !reflect.DeepEqual(i.Object, s) {
		t.Fatalf("expected %+v. got=%+v", s, i.Object)
	}
}

func TestIntersectionAggregation(t *testing.T) {
	s := Sphere()
	i1 := Intersection{s, 1}
	i2 := Intersection{s, 2}
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

func TestIntersectionHit(t *testing.T) {
	s := Sphere()
	i1 := Intersection{s, 1}
	i2 := Intersection{s, 2}
	xs := Intersections{i2, i1}

	i, _ := xs.Hit()
	if !reflect.DeepEqual(i, i1) {
		t.Fatalf("expected %+v. got=%+v", i1, i)
	}

	s = Sphere()
	i1 = Intersection{s, -1}
	i2 = Intersection{s, 1}
	xs = Intersections{i2, i1}
	i, _ = xs.Hit()
	if !reflect.DeepEqual(i, i2) {
		t.Fatalf("expected %+v. got=%+v", i2, i)
	}

	s = Sphere()
	i1 = Intersection{s, -2}
	i2 = Intersection{s, -1}
	xs = Intersections{i2, i1}
	_, hit := xs.Hit()
	if hit {
		t.Fatal("expected no hit")
	}

	s = Sphere()
	i1 = Intersection{s, 5}
	i2 = Intersection{s, 7}
	i3 := Intersection{s, -3}
	i4 := Intersection{s, 2}
	xs = Intersections{i1, i2, i3, i4}
	i, _ = xs.Hit()
	if !reflect.DeepEqual(i, i4) {
		t.Fatalf("expected %+v. got=%+v", i4, i)
	}
}

func TestIntersectionPrecompute(t *testing.T) {
	r := Ray{Point{Z: -5}, Vector{Z: 1}}
	shape := Sphere()
	i := Intersection{shape, 4}
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
	r := Ray{Point{}, Vector{Z: 1}}
	shape := Sphere()
	i := Intersection{shape, 1}
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

func TestHitPointOffset(t *testing.T) {
	r := Ray{Point{Z: -5}, Vector{Z: 1}}
	shape := Sphere()
	shape.Transform = Translation(0, 0, 1)
	i := Intersection{shape, 5}
	comps := i.PrepareComputations(r)

	if comps.OverPoint.Z >= -Epsilon/2 {
		t.Fatalf("expected %f. got=%f", -Epsilon/2, comps.OverPoint.Z)
	}
	if comps.Point.Z < comps.OverPoint.Z {
		t.Fatalf("expected %f. got=%f", comps.OverPoint.Z, comps.Point.Z)
	}
}
