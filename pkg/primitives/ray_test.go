package primitives

import (
	"reflect"
	"testing"
)

func TestRayCreation(t *testing.T) {
	origin := Point{XVal: 1, YVal: 2, ZVal: 3}
	direction := Vector{XVal: 4, YVal: 5, ZVal: 6}
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
		Point{XVal: 2, YVal: 3, ZVal: 4},
		Vector{XVal: 1, YVal: 0, ZVal: 0},
	}

	tests := map[float64]Point{
		0:   {XVal: 2, YVal: 3, ZVal: 4},
		1:   {XVal: 3, YVal: 3, ZVal: 4},
		-1:  {XVal: 1, YVal: 3, ZVal: 4},
		2.5: {XVal: 4.5, YVal: 3, ZVal: 4},
	}

	for time, expected := range tests {
		actual, err := Position(r, time)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(actual, &expected) {
			t.Fatalf("expected %+v. got=%+v", &expected, actual)
		}
	}
}
