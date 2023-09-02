package shapes

import (
	"math"

	"github.com/adriffaud/ray-tracer-challenge/pkg/primitives"
)

type Shape struct{}

func (s Shape) Intersect(r primitives.Ray) Intersections {
	s2r := r.Origin.SubPoint(primitives.Point{X: 0, Y: 0, Z: 0})
	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(s2r)
	c := s2r.Dot(s2r) - 1

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return Intersections{}
	}

	return Intersections{
		{T: (-b - math.Sqrt(discriminant)) / (2 * a), Object: s},
		{T: (-b + math.Sqrt(discriminant)) / (2 * a), Object: s},
	}
}
