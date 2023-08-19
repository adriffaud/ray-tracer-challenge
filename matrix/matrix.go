package matrix

import (
	"errors"

	"github.com/adriffaud/ray-tracer-challenge/tuple"
)

type Matrix [][]float64

// NewMatrix generates a Matrix with the given x/y size.
func NewMatrix(x, y int) Matrix {
	m := make(Matrix, x)
	for row := range m {
		m[row] = make([]float64, y)
	}

	return m
}

func IdentityMatrix() Matrix {
	return Matrix{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

func Multiply(a, b Matrix) Matrix {
	m := NewMatrix(4, 4)

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			m[row][col] = a[row][0]*b[0][col] + a[row][1]*b[1][col] + a[row][2]*b[2][col] + a[row][3]*b[3][col]
		}
	}

	return m
}

func MultiplyTuple(m Matrix, t tuple.Tuple) (tuple.Tuple, error) {
	switch t.W() {
	case 0:
		v := tuple.Vector{
			XVal: m[0][0]*t.X() + m[0][1]*t.Y() + m[0][2]*t.Z(),
			YVal: m[1][0]*t.X() + m[1][1]*t.Y() + m[1][2]*t.Z(),
			ZVal: m[2][0]*t.X() + m[2][1]*t.Y() + m[2][2]*t.Z(),
		}

		return &v, nil
	case 1:
		p := tuple.Point{
			XVal: m[0][0]*t.X() + m[0][1]*t.Y() + m[0][2]*t.Z() + m[0][3],
			YVal: m[1][0]*t.X() + m[1][1]*t.Y() + m[1][2]*t.Z() + m[1][3],
			ZVal: m[2][0]*t.X() + m[2][1]*t.Y() + m[2][2]*t.Z() + m[2][3],
		}

		return &p, nil
	default:
		return nil, errors.New("operation not allowed")
	}
}

func Transpose(m Matrix) Matrix {
	t := NewMatrix(4, 4)

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			t[col][row] = m[row][col]
		}
	}

	return t
}
