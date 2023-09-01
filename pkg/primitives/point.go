package primitives

// Point is a struct representing a 3D point
type Point struct{ XVal, YVal, ZVal float64 }

// X returns the x-coordinate of the Point
func (p *Point) X() float64 {
	return p.XVal
}

// Y returns the y-coordinate of the Point
func (p *Point) Y() float64 {
	return p.YVal
}

// Z returns the z-coordinate of the Point
func (p *Point) Z() float64 {
	return p.ZVal
}

// W returns the w-coordinate of the Point
func (p *Point) W() float64 {
	return 1.0
}
