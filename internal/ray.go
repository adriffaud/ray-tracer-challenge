package internal

type Ray struct {
	Origin    Point
	Direction Vector
}

func (r Ray) Position(t float64) Point {
	multiplied := r.Direction.Multiply(t)
	return r.Origin.Add(multiplied)
}

func (r Ray) Transform(m Matrix) Ray {
	return Ray{
		Origin:    r.Origin.MultiplyMatrix(m),
		Direction: r.Direction.MultiplyMatrix(m),
	}
}
