package matrix

func Translation(x, y, z float64) Matrix {
	m := IdentityMatrix()
	m[0][3] = x
	m[1][3] = y
	m[2][3] = z
	return m
}
