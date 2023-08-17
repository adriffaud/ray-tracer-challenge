package color

type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

func NewColor(r, g, b float64) *Color {
	return &Color{Red: r, Green: g, Blue: b}
}

func Add(c1, c2 *Color) *Color {
	return &Color{Red: c1.Red + c2.Red, Green: c1.Green + c2.Green, Blue: c1.Blue + c2.Blue}
}

func Sub(c1, c2 *Color) *Color {
	return &Color{Red: c1.Red - c2.Red, Green: c1.Green - c2.Green, Blue: c1.Blue - c2.Blue}
}

func MultiplyByScalar(c *Color, s float64) *Color {
	return &Color{Red: c.Red * s, Green: c.Green * s, Blue: c.Blue * s}
}

func Multiply(c1, c2 *Color) *Color {
	return &Color{Red: c1.Red * c2.Red, Green: c1.Green * c2.Green, Blue: c1.Blue * c2.Blue}
}
