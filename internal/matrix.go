package internal

import "errors"

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

func (a Matrix) Multiply(b Matrix) Matrix {
	m := NewMatrix(4)

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			m[row][col] = a[row][0]*b[0][col] + a[row][1]*b[1][col] + a[row][2]*b[2][col] + a[row][3]*b[3][col]
		}
	}

	return m
}

func (m Matrix) Transpose() Matrix {
	t := NewMatrix(4)

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			t[col][row] = m[row][col]
		}
	}

	return t
}

func (m Matrix) Determinant() float64 {
	var det float64

	if len(m) == 2 {
		det = m[0][0]*m[1][1] - m[0][1]*m[1][0]
	} else {
		for col := 0; col < len(m); col++ {
			det = det + m[0][col]*m.Cofactor(0, col)
		}
	}

	return det
}

func (m Matrix) Submatrix(row, col int) Matrix {
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

func (m Matrix) Minor(row, col int) float64 {
	sub := m.Submatrix(row, col)
	return sub.Determinant()
}

func (m Matrix) Cofactor(row, col int) float64 {
	minor := m.Minor(row, col)

	if (row+col)%2 == 1 {
		return -minor
	} else {
		return minor
	}
}

func (m Matrix) Inverse() (Matrix, error) {
	det := m.Determinant()
	if det == 0 {
		return nil, errors.New("matrix is not invertible")
	}

	res := NewMatrix(len(m))

	for row := 0; row < len(m); row++ {
		for col := 0; col < len(m); col++ {
			c := m.Cofactor(row, col)
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
			if !ApproxEq(a[row][col], b[row][col]) {
				return false
			}
		}
	}

	return true
}
