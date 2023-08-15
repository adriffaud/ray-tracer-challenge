package tuple

// Vector is a struct representing a 3D vector
type Vector struct{ XVal, YVal, ZVal float32 }

// NewVector creates a new Vector instance
func NewVector(x, y, z float32) *Vector {
	return &Vector{XVal: x, YVal: y, ZVal: z}
}

// X returns the x-coordinate of the Vector
func (v *Vector) X() float32 {
	return v.XVal
}

// Y returns the y-coordinate of the Vector
func (v *Vector) Y() float32 {
	return v.YVal
}

// Z returns the z-coordinate of the Vector
func (v *Vector) Z() float32 {
	return v.ZVal
}

// W returns the w-coordinate of the Vector
func (v *Vector) W() float32 {
	return 0.0
}
