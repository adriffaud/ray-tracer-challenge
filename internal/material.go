package internal

type Material struct {
	Color                                 Color
	Ambient, Diffuse, Specular, Shininess float64
}

func NewMaterial() Material {
	return Material{
		Color:     Color{R: 1, G: 1, B: 1},
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
}
