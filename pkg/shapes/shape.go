package shapes

import (
	"fmt"
	"math"

	"github.com/adriffaud/ray-tracer-challenge/pkg/primitives"
)

type Shape struct{}

func (s Shape) Intersect(r primitives.Ray) Intersections {
	s2r, err := primitives.Sub(&r.Origin, &primitives.Point{XVal: 0, YVal: 0, ZVal: 0})
	if err != nil {
		fmt.Print(err)
	}

	a := primitives.Dot(&r.Direction, &r.Direction)
	b := 2 * primitives.Dot(&r.Direction, s2r)
	c := primitives.Dot(s2r, s2r) - 1

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return Intersections{}
	}

	return Intersections{
		{T: (-b - math.Sqrt(discriminant)) / (2 * a), Object: s},
		{T: (-b + math.Sqrt(discriminant)) / (2 * a), Object: s},
	}
}
