package internal

import (
	"reflect"
	"testing"
)

func TestRayCreation(t *testing.T) {
	origin := Point{X: 1, Y: 2, Z: 3}
	direction := Vector{X: 4, Y: 5, Z: 6}
	r := Ray{Origin: origin, Direction: direction}

	if r.Origin != origin {
		t.Fatalf("expected origin to be %+v. got=%+v", origin, r.Origin)
	}
	if r.Direction != direction {
		t.Fatalf("expected origin to be %+v. got=%+v", direction, r.Direction)
	}
}

func TestComputingPointFromDistance(t *testing.T) {
	r := Ray{
		Point{X: 2, Y: 3, Z: 4},
		Vector{X: 1},
	}

	tests := map[float64]Point{
		0:   {X: 2, Y: 3, Z: 4},
		1:   {X: 3, Y: 3, Z: 4},
		-1:  {X: 1, Y: 3, Z: 4},
		2.5: {X: 4.5, Y: 3, Z: 4},
	}

	for time, expected := range tests {
		actual := r.Position(time)

		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("expected %+v. got=%+v", expected, actual)
		}
	}
}

func TestRayTranslation(t *testing.T) {
	r := Ray{
		Origin:    Point{X: 1, Y: 2, Z: 3},
		Direction: Vector{Y: 1},
	}
	m := Translation(3, 4, 5)
	r2 := r.Transform(m)
	expected_origin := Point{X: 4, Y: 6, Z: 8}

	if !reflect.DeepEqual(r2.Origin, expected_origin) {
		t.Fatalf("expected %+v. got=%+v", expected_origin, r2.Origin)
	}
	if !reflect.DeepEqual(r2.Direction, r.Direction) {
		t.Fatalf("expected %+v. got=%+v", r.Direction, r2.Direction)
	}
}

func TestRayScaling(t *testing.T) {
	r := Ray{
		Origin:    Point{X: 1, Y: 2, Z: 3},
		Direction: Vector{Y: 1},
	}
	m := Scaling(2, 3, 4)
	r2 := r.Transform(m)
	expected_origin := Point{X: 2, Y: 6, Z: 12}
	expected_direction := Vector{Y: 3}

	if !reflect.DeepEqual(r2.Origin, expected_origin) {
		t.Fatalf("expected %+v. got=%+v", expected_origin, r2.Origin)
	}
	if !reflect.DeepEqual(r2.Direction, expected_direction) {
		t.Fatalf("expected %+v. got=%+v", expected_direction, r2.Direction)
	}
}
