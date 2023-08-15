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

func Magnitude(v Vector) float64 {
	return math.Sqrt(v.XVal*v.XVal + v.YVal*v.YVal + v.ZVal*v.ZVal)
}

func Normalize(v Vector) Vector {
	magnitude := Magnitude(v)
	return *NewVector(v.XVal/magnitude, v.YVal/magnitude, v.ZVal/magnitude)
}
