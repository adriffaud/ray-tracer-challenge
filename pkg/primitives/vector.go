package primitives

// Vector is a struct representing a 3D vector
type Vector struct{ XVal, YVal, ZVal float64 }

// X returns the x-coordinate of the Vector
func (v *Vector) X() float64 {
	return v.XVal
}

// Y returns the y-coordinate of the Vector
func (v *Vector) Y() float64 {
	return v.YVal
}

// Z returns the z-coordinate of the Vector
func (v *Vector) Z() float64 {
	return v.ZVal
}

// W returns the w-coordinate of the Vector
func (v *Vector) W() float64 {
	return 0.0
}
