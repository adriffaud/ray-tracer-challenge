package internal

// Point is a struct representing a 3D point
type Point struct{ X, Y, Z float64 }

// Add adds two tuples and returns a new tuple
func (a Point) Add(b Vector) Point {
	return Point{X: a.X + b.X, Y: a.Y + b.Y, Z: a.Z + b.Z}
}

func (a Point) SubPoint(b Point) Vector {
	return Vector{X: a.X - b.X, Y: a.Y - b.Y, Z: a.Z - b.Z}
}

func (a Point) SubVector(b Vector) Point {
	return Point{X: a.X - b.X, Y: a.Y - b.Y, Z: a.Z - b.Z}
}

func (p Point) Negate() Point {
	return Point{X: -p.X, Y: -p.Y, Z: -p.Z}
}

func (p Point) Multiply(s float64) Point {
	return Point{X: p.X * s, Y: p.Y * s, Z: p.Z * s}
}

func (p Point) Divide(s float64) Point {
	return Point{X: p.X / s, Y: p.Y / s, Z: p.Z / s}
}

func (p Point) MultiplyMatrix(m Matrix) Point {
	return Point{
		X: m[0][0]*p.X + m[0][1]*p.Y + m[0][2]*p.Z + m[0][3],
		Y: m[1][0]*p.X + m[1][1]*p.Y + m[1][2]*p.Z + m[1][3],
		Z: m[2][0]*p.X + m[2][1]*p.Y + m[2][2]*p.Z + m[2][3],
	}
}
