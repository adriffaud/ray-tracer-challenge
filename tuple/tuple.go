package tuple

import "errors"

type Tuple interface {
	X() float32
	Y() float32
	Z() float32
	W() float32
}

// Add adds two tuples and returns a new tuple
func Add(a, b Tuple) (Tuple, error) {
	x := a.X() + b.X()
	y := a.Y() + b.Y()
	z := a.Z() + b.Z()
	w := a.W() + b.W()

	switch w {
	case 1:
		return NewPoint(x, y, z), nil
	case 0:
		return NewVector(x, y, z), nil
	default:
		return nil, errors.New("operation not allowed")
	}
}

// Sub subtracts two tuples and returns a new tuple
func Sub(a, b Tuple) (Tuple, error) {
	x := a.X() - b.X()
	y := a.Y() - b.Y()
	z := a.Z() - b.Z()
	w := a.W() - b.W()

	switch w {
	case 1:
		return NewPoint(x, y, z), nil
	case 0:
		return NewVector(x, y, z), nil
	default:
		return nil, errors.New("operation not allowed")
	}
}
