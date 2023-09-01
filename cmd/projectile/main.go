package main

import (
	"github.com/adriffaud/ray-tracer-challenge/pkg/canvas"
	"github.com/adriffaud/ray-tracer-challenge/pkg/color"
	"github.com/adriffaud/ray-tracer-challenge/pkg/primitives"
)

type Projectile struct {
	Position primitives.Tuple
	Velocity primitives.Tuple
}

type Environment struct {
	Gravity primitives.Vector
	Wind    primitives.Vector
}

func (p *Projectile) Tick(env Environment) {
	position, err := primitives.Add(p.Position, p.Velocity)
	if err != nil {
		panic(err)
	}
	p.Position = position

	intermediate, err := primitives.Add(p.Velocity, &env.Gravity)
	if err != nil {
		panic(err)
	}
	velocity, err := primitives.Add(intermediate, &env.Wind)
	if err != nil {
		panic(err)
	}
	p.Velocity = velocity
}

func main() {
	velocity, err := primitives.Multiply(primitives.Normalize(&primitives.Vector{XVal: 1, YVal: 1.8, ZVal: 0}), 11.25)
	if err != nil {
		panic(err)
	}

	p := Projectile{
		Position: &primitives.Point{XVal: 0, YVal: 1, ZVal: 0},
		Velocity: velocity,
	}

	env := Environment{
		Gravity: primitives.Vector{XVal: 0, YVal: -0.1, ZVal: 0},
		Wind:    primitives.Vector{XVal: -0.01, YVal: 0, ZVal: 0},
	}

	c := canvas.NewCanvas(900, 550)

	for p.Position.Y() > 0 {
		p.Tick(env)
		row := int(500 - p.Position.Y() - 1)
		col := int(p.Position.X())
		c.WritePixel(col, row, color.Color{R: 1})
	}

	canvas.Export(c, "projectile.png")
}
