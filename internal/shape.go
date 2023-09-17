package internal

import "math"

type Shape struct {
	Transform Matrix
	Material  Material
}

func (s Shape) Intersect(r Ray) Intersections {
	inv, err := s.Transform.Inverse()
	if err != nil {
		panic(err)
	}
	r2 := r.Transform(inv)

	s2r := r2.Origin.SubPoint(Point{X: 0, Y: 0, Z: 0})
	a := r2.Direction.Dot(r2.Direction)
	b := 2 * r2.Direction.Dot(s2r)
	c := s2r.Dot(s2r) - 1

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return Intersections{}
	}

	return Intersections{
		{Distance: (-b - math.Sqrt(discriminant)) / (2 * a), Object: s},
		{Distance: (-b + math.Sqrt(discriminant)) / (2 * a), Object: s},
	}
}

func (s Shape) NormalAt(p Point) Vector {
	inv, err := s.Transform.Inverse()
	if err != nil {
		panic(err)
	}

	objPoint := p.MultiplyMatrix(inv)
	objNormal := objPoint.SubPoint(Point{})
	worldNormal := objNormal.MultiplyMatrix(inv.Transpose())

	return worldNormal.Normalize()
}
