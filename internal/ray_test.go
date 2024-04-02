package internal

import (
	"testing"
)

func TestRayCreation(t *testing.T) {
	origin := Point{X: 1, Y: 2, Z: 3}
	direction := Vector{X: 4, Y: 5, Z: 6}
	r := Ray{Origin: origin, Direction: direction}

	assertEqual(t, origin, r.Origin)
	assertEqual(t, direction, r.Direction)
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

		assertDeepEqual(t, expected, actual)
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

	assertDeepEqual(t, r2.Origin, expected_origin)
	assertDeepEqual(t, r2.Direction, r.Direction)
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

	assertDeepEqual(t, r2.Origin, expected_origin)
	assertDeepEqual(t, r2.Direction, expected_direction)
}
