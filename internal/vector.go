package internal

import "math"

// Vector is a struct representing a 3D vector
type Vector struct{ X, Y, Z float64 }

// Add adds two vectors and returns a new vector
func (a Vector) Add(b Vector) Vector {
	return Vector{X: a.X + b.X, Y: a.Y + b.Y, Z: a.Z + b.Z}
}

func (a Vector) Sub(b Vector) Vector {
	return Vector{X: a.X - b.X, Y: a.Y - b.Y, Z: a.Z - b.Z}
}

func (v Vector) Negate() Vector {
	return Vector{X: -v.X, Y: -v.Y, Z: -v.Z}
}

func (v Vector) Multiply(s float64) Vector {
	return Vector{X: v.X * s, Y: v.Y * s, Z: v.Z * s}
}

func (v Vector) Divide(s float64) Vector {
	return Vector{X: v.X / s, Y: v.Y / s, Z: v.Z / s}
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector) Normalize() Vector {
	magnitude := v.Magnitude()
	return Vector{X: v.X / magnitude, Y: v.Y / magnitude, Z: v.Z / magnitude}
}

func (v1 Vector) Dot(v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func (v1 Vector) Cross(v2 Vector) Vector {
	return Vector{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: v1.Z*v2.X - v1.X*v2.Z,
		Z: v1.X*v2.Y - v1.Y*v2.X,
	}
}

func (v Vector) MultiplyMatrix(m Matrix) Vector {
	return Vector{
		X: m[0][0]*v.X + m[0][1]*v.Y + m[0][2]*v.Z,
		Y: m[1][0]*v.X + m[1][1]*v.Y + m[1][2]*v.Z,
		Z: m[2][0]*v.X + m[2][1]*v.Y + m[2][2]*v.Z,
	}
}

func (v Vector) Reflect(n Vector) Vector {
	return v.Sub(n.Multiply(2 * v.Dot(n)))
}
