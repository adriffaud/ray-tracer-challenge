package main

import (
	"github.com/adriffaud/ray-tracer-challenge/pkg/canvas"
	"github.com/adriffaud/ray-tracer-challenge/pkg/color"
	"github.com/adriffaud/ray-tracer-challenge/pkg/primitives"
)

type Projectile struct {
	Position primitives.Point
	Velocity primitives.Vector
}

type Environment struct {
	Gravity primitives.Vector
	Wind    primitives.Vector
}

func (p *Projectile) Tick(env Environment) {
	position := p.Position.Add(p.Velocity)
	p.Position = position

	intermediate := p.Velocity.Add(env.Gravity)
	velocity := intermediate.Add(env.Wind)
	p.Velocity = velocity
}

func main() {
	velocity := primitives.Vector{X: 1, Y: 1.8, Z: 0}.Normalize().Multiply(11.25)

	p := Projectile{
		Position: primitives.Point{X: 0, Y: 1, Z: 0},
		Velocity: velocity,
	}

	env := Environment{
		Gravity: primitives.Vector{X: 0, Y: -0.1, Z: 0},
		Wind:    primitives.Vector{X: -0.01, Y: 0, Z: 0},
	}

	c := canvas.NewCanvas(900, 550)

	for p.Position.Y > 0 {
		p.Tick(env)
		row := int(500 - p.Position.Y - 1)
		col := int(p.Position.X)
		c.WritePixel(col, row, color.Color{R: 1})
	}

	canvas.Export(c, "projectile.png")
}
