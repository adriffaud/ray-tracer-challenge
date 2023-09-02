package primitives

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
