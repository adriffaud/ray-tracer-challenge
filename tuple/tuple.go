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
		return &Vector{x, y, z}, nil
	case 1:
		return &Point{x, y, z}, nil
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
		return &Vector{x, y, z}, nil
	case 1:
		return &Point{x, y, z}, nil
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
		return &Vector{x, y, z}, nil
	case 1:
		return &Point{x, y, z}, nil
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
		return &Vector{x, y, z}, nil
	case 1:
		return &Point{x, y, z}, nil
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
		return &Vector{x, y, z}, nil
	case 1:
		return &Point{x, y, z}, nil
	default:
		return nil, errors.New("operation not allowed")
	}
}

func Magnitude(t Tuple) float64 {
	return math.Sqrt(t.X()*t.X() + t.Y()*t.Y() + t.Z()*t.Z())
}

func Normalize(t Tuple) Tuple {
	magnitude := Magnitude(t)
	return &Vector{XVal: t.X() / magnitude, YVal: t.Y() / magnitude, ZVal: t.Z() / magnitude}
}

func Dot(v1, v2 Tuple) float64 {
	return v1.X()*v2.X() + v1.Y()*v2.Y() + v1.Z()*v2.Z()
}

func Cross(v1, v2 Tuple) Tuple {
	x := v1.Y()*v2.Z() - v1.Z()*v2.Y()
	y := v1.Z()*v2.X() - v1.X()*v2.Z()
	z := v1.X()*v2.Y() - v1.Y()*v2.X()

	return &Vector{x, y, z}
}
