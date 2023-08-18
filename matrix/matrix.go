package matrix

import "github.com/adriffaud/ray-tracer-challenge/tuple"

type Matrix [][]float64

// NewMatrix generates a Matrix with the given x/y size.
func NewMatrix(x, y int) Matrix {
	m := make(Matrix, x)
	for row := range m {
		m[row] = make([]float64, y)
	}

	return m
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

func MultiplyTuple(m Matrix, t tuple.Tuple) tuple.Tuple {
	return &tuple.Point{}
}
