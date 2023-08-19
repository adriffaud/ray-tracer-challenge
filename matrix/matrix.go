package matrix

import (
	"errors"

	"github.com/adriffaud/ray-tracer-challenge/float"
	"github.com/adriffaud/ray-tracer-challenge/tuple"
)

type Matrix [][]float64

// NewMatrix generates a Matrix with the given x/y size.
func NewMatrix(size int) Matrix {
	m := make(Matrix, size)
	for row := range m {
		m[row] = make([]float64, size)
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
	m := NewMatrix(4)

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
	t := NewMatrix(4)

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			t[col][row] = m[row][col]
		}
	}

	return t
}

func Determinant(m Matrix) int {
	var det int

	if len(m) == 2 {
		det = int(m[0][0]*m[1][1] - m[0][1]*m[1][0])
	} else {
		for col := 0; col < len(m); col++ {
			det = det + int(m[0][col])*Cofactor(m, 0, col)
		}
	}

	return det
}

func Submatrix(m Matrix, row, col int) Matrix {
	res := NewMatrix(len(m) - 1)

	rowCounter := 0
	for i := 0; i < len(m); i++ {
		if i == row {
			continue
		}

		colCounter := 0
		for j := 0; j < len(m); j++ {
			if j == col {
				continue
			}

			res[rowCounter][colCounter] = m[i][j]

			colCounter += 1
		}

		rowCounter += 1
	}

	return res
}

func Minor(m Matrix, row, col int) int {
	sub := Submatrix(m, row, col)
	return Determinant(sub)
}

func Cofactor(m Matrix, row, col int) int {
	minor := Minor(m, row, col)

	if (row+col)%2 == 1 {
		return -minor
	} else {
		return minor
	}
}

func Inverse(m Matrix) (Matrix, error) {
	det := Determinant(m)
	if det == 0 {
		return nil, errors.New("matrix is not invertible")
	}

	res := NewMatrix(len(m))

	for row := 0; row < len(m); row++ {
		for col := 0; col < len(m); col++ {
			c := Cofactor(m, row, col)
			res[col][row] = float64(c) / float64(det)
		}
	}

	return res, nil
}

func Eq(a, b Matrix) bool {
	if len(a) != len(b) {
		return false
	}

	for row := 0; row < len(a); row++ {
		for col := 0; col < len(a); col++ {
			if !float.ApproxEq(a[row][col], b[row][col]) {
				return false
			}
		}
	}

	return true
}
