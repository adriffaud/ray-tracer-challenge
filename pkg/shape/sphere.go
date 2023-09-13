package shape

import "github.com/adriffaud/ray-tracer-challenge/pkg/primitives"

func Sphere() Shape {
	return Shape{
		Transform: primitives.IdentityMatrix(),
		Material:  primitives.NewMaterial(),
	}
}
