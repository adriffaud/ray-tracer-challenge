package internal

type Color struct {
	R float64
	G float64
	B float64
}

func (c1 Color) Add(c2 Color) Color {
	return Color{R: c1.R + c2.R, G: c1.G + c2.G, B: c1.B + c2.B}
}

func (c1 Color) Sub(c2 Color) Color {
	return Color{R: c1.R - c2.R, G: c1.G - c2.G, B: c1.B - c2.B}
}

func (c Color) MultiplyScalar(s float64) Color {
	return Color{R: c.R * s, G: c.G * s, B: c.B * s}
}

func (c1 Color) Multiply(c2 Color) Color {
	return Color{R: c1.R * c2.R, G: c1.G * c2.G, B: c1.B * c2.B}
}
