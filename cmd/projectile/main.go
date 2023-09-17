package main

import "github.com/adriffaud/ray-tracer-challenge/internal"

type Projectile struct {
	Position internal.Point
	Velocity internal.Vector
}

type Environment struct {
	Gravity internal.Vector
	Wind    internal.Vector
}

func (p *Projectile) Tick(env Environment) {
	position := p.Position.Add(p.Velocity)
	p.Position = position

	intermediate := p.Velocity.Add(env.Gravity)
	velocity := intermediate.Add(env.Wind)
	p.Velocity = velocity
}

func main() {
	velocity := internal.Vector{X: 1, Y: 1.8, Z: 0}.Normalize().Multiply(11.25)

	p := Projectile{
		Position: internal.Point{X: 0, Y: 1, Z: 0},
		Velocity: velocity,
	}

	env := Environment{
		Gravity: internal.Vector{X: 0, Y: -0.1, Z: 0},
		Wind:    internal.Vector{X: -0.01, Y: 0, Z: 0},
	}

	c := internal.NewCanvas(900, 550)

	for p.Position.Y > 0 {
		p.Tick(env)
		row := int(500 - p.Position.Y - 1)
		col := int(p.Position.X)
		c.WritePixel(col, row, internal.Color{R: 1})
	}

	internal.Export(c, "projectile.png")
}
