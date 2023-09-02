package primitives

type Ray struct {
	Origin    Point
	Direction Vector
}

func (r Ray) Position(t float64) Point {
	multiplied := r.Direction.Multiply(t)
	return r.Origin.Add(multiplied)
}
}
