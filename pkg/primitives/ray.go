package primitives

type Ray struct {
	Origin    Point
	Direction Vector
}

func Position(r Ray, t float64) (Tuple, error) {
	multiplied, err := Multiply(&r.Direction, t)
	if err != nil {
		return nil, err
	}
	return Add(&r.Origin, multiplied)
}
