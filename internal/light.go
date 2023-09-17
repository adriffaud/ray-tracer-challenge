package internal

import "math"

type Light struct {
	Intensity Color
	Position  Point
}

func Lighting(m Material, l Light, p Point, eye, normal Vector) Color {
	effective := m.Color.Multiply(l.Intensity)
	lightv := l.Position.SubPoint(p).Normalize()
	ambient := effective.MultiplyScalar(m.Ambient)
	lightDotNormal := lightv.Dot(normal)

	var diffuse, specular Color

	if lightDotNormal >= 0 {
		diffuse = effective.MultiplyScalar(m.Diffuse * lightDotNormal)
		reflect := lightv.Negate().Reflect(normal)
		reflectDotEye := reflect.Dot(eye)

		if reflectDotEye > 0 {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = l.Intensity.MultiplyScalar(m.Specular * factor)
		}
	}

	return ambient.Add(diffuse).Add(specular)
}
