package tuple

// Point is a struct representing a 3D point
type Point struct{ XVal, YVal, ZVal float32 }

// NewPoint creates a new Point instance
func NewPoint(x, y, z float32) *Point {
	return &Point{XVal: x, YVal: y, ZVal: z}
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
