package tuple

import (
	"errors"
	"math"
)

type Tuple interface {
	X() float64
	Y() float64
	Z() float64
	W() float64
}

type Negatable interface {
	Negate() Tuple
}

// Add adds two tuples and returns a new tuple
func Add(a, b Tuple) (Tuple, error) {
	x := a.X() + b.X()
	y := a.Y() + b.Y()
	z := a.Z() + b.Z()
	w := a.W() + b.W()

	switch w {
	case 0:
		return NewVector(x, y, z), nil
	case 1:
		return NewPoint(x, y, z), nil
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
	case 0:
		return NewVector(x, y, z), nil
	case 1:
		return NewPoint(x, y, z), nil
	default:
		return nil, errors.New("operation not allowed")
	}
}

// Negate return the a new Tuple instance with its negated values
func Negate(t Tuple) (Tuple, error) {
	x := -t.X()
	y := -t.Y()
	z := -t.Z()

	switch t.W() {
	case 0:
		return NewVector(x, y, z), nil
	case 1:
		return NewPoint(x, y, z), nil
	default:
		return nil, errors.New("operation not allowed")
	}
}

func Multiply(t Tuple, s float64) (Tuple, error) {
	x := t.X() * s
	y := t.Y() * s
	z := t.Z() * s

	switch t.W() {
	case 0:
		return NewVector(x, y, z), nil
	case 1:
		return NewPoint(x, y, z), nil
	default:
		return nil, errors.New("operation not allowed")
	}
}

func Divide(t Tuple, s float64) (Tuple, error) {
	x := t.X() / s
	y := t.Y() / s
	z := t.Z() / s

	switch t.W() {
	case 0:
		return NewVector(x, y, z), nil
	case 1:
		return NewPoint(x, y, z), nil
	default:
		return nil, errors.New("operation not allowed")
	}
}

func Magnitude(v Tuple) float64 {
	return math.Sqrt(v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z())
}

func Normalize(v Tuple) Vector {
	magnitude := Magnitude(v)
	return *NewVector(v.X()/magnitude, v.Y()/magnitude, v.Z()/magnitude)
}

func Dot(v1, v2 Tuple) float64 {
	return v1.X()*v2.X() + v1.Y()*v2.Y() + v1.Z()*v2.Z()
}
