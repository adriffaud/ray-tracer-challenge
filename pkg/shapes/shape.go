package shapes

import (
	"math"

	"github.com/adriffaud/ray-tracer-challenge/pkg/primitives"
)

type Shape struct {
	Transform primitives.Matrix
}

func (s Shape) Intersect(r primitives.Ray) Intersections {
	inv, err := s.Transform.Inverse()
	if err != nil {
		panic(err)
	}
	r2 := r.Transform(inv)

	s2r := r2.Origin.SubPoint(primitives.Point{X: 0, Y: 0, Z: 0})
	a := r2.Direction.Dot(r2.Direction)
	b := 2 * r2.Direction.Dot(s2r)
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

func (s Shape) NormalAt(p primitives.Point) primitives.Vector {
	inv, err := s.Transform.Inverse()
	if err != nil {
		panic(err)
	}

	objPoint := p.MultiplyMatrix(inv)
	objNormal := objPoint.SubPoint(primitives.Point{})
	worldNormal := objNormal.MultiplyMatrix(inv.Transpose())

	return worldNormal.Normalize()
}
