package shapes

import (
	"github.com/adriffaud/ray-tracer-challenge/pkg/component"
	"github.com/adriffaud/ray-tracer-challenge/pkg/primitives"
)

func Sphere() Shape {
	return Shape{
		Transform: primitives.IdentityMatrix(),
		Material:  component.NewMaterial(),
	}
}
