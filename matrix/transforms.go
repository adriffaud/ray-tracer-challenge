package matrix

import (
	"math"
)

func Translation(x, y, z float64) Matrix {
	m := IdentityMatrix()
	m[0][3] = x
	m[1][3] = y
	m[2][3] = z
	return m
}

func Scaling(x, y, z float64) Matrix {
	m := IdentityMatrix()
	m[0][0] = x
	m[1][1] = y
	m[2][2] = z
	return m
}

func RotationX(r float64) Matrix {
	m := IdentityMatrix()
	m[1][1] = math.Cos(r)
	m[1][2] = -math.Sin(r)
	m[2][1] = math.Sin(r)
	m[2][2] = math.Cos(r)

	return m
}

func RotationY(r float64) Matrix {
	m := IdentityMatrix()
	m[0][0] = math.Cos(r)
	m[0][2] = math.Sin(r)
	m[2][0] = -math.Sin(r)
	m[2][2] = math.Cos(r)

	return m
}

func RotationZ(r float64) Matrix {
	m := IdentityMatrix()
	m[0][0] = math.Cos(r)
	m[0][1] = -math.Sin(r)
	m[1][0] = math.Sin(r)
	m[1][1] = math.Cos(r)

	return m
}
