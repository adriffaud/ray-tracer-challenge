package color

type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

func NewColor(r, g, b float64) *Color { return &Color{Red: r, Green: g, Blue: b} }
