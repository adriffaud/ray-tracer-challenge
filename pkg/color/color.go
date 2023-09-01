package color

type Color struct {
	R float64
	G float64
	B float64
}

func Add(c1, c2 Color) Color {
	return Color{R: c1.R + c2.R, G: c1.G + c2.G, B: c1.B + c2.B}
}

func Sub(c1, c2 Color) Color {
	return Color{R: c1.R - c2.R, G: c1.G - c2.G, B: c1.B - c2.B}
}

func MultiplyByScalar(c Color, s float64) Color {
	return Color{R: c.R * s, G: c.G * s, B: c.B * s}
}

func Multiply(c1, c2 Color) Color {
	return Color{R: c1.R * c2.R, G: c1.G * c2.G, B: c1.B * c2.B}
}
