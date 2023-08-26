package matrix

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
