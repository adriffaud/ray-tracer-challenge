package tuple

type Tuple interface {
	X() float32
	Y() float32
	Z() float32
	W() float32
}

// Point is a struct representing a 3D point
type Point struct{ XVal, YVal, ZVal float32 }

// Vector is a struct representing a 3D vector
type Vector struct{ XVal, YVal, ZVal float32 }

func NewPoint(x, y, z float32) *Point {
	return &Point{XVal: x, YVal: y, ZVal: z}
}

// NewVector creates a new Vector instance
func NewVector(x, y, z float32) *Vector {
	return &Vector{XVal: x, YVal: y, ZVal: z}
}

// X returns the x-coordinate of the Point
func (p *Point) X() float32 {
	return p.XVal
}

// Y returns the y-coordinate of the Point
func (p *Point) Y() float32 {
	return p.YVal
}

// Z returns the z-coordinate of the Point
func (p *Point) Z() float32 {
	return p.ZVal
}

// W returns the w-coordinate of the Point
func (p *Point) W() float32 {
	return 1.0
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
