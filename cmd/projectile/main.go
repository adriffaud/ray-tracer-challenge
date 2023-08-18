package main

import (
	"github.com/adriffaud/ray-tracer-challenge/canvas"
	"github.com/adriffaud/ray-tracer-challenge/color"
	"github.com/adriffaud/ray-tracer-challenge/tuple"
)

type Projectile struct {
	Position tuple.Tuple
	Velocity tuple.Tuple
}

type Environment struct {
	Gravity tuple.Vector
	Wind    tuple.Vector
}

func (p *Projectile) Tick(env Environment) {
	position, err := tuple.Add(p.Position, p.Velocity)
	if err != nil {
		panic(err)
	}
	p.Position = position

	intermediate, err := tuple.Add(p.Velocity, &env.Gravity)
	if err != nil {
		panic(err)
	}
	velocity, err := tuple.Add(intermediate, &env.Wind)
	if err != nil {
		panic(err)
	}
	p.Velocity = velocity
}

func main() {
	velocity, err := tuple.Multiply(tuple.Normalize(&tuple.Vector{XVal: 1, YVal: 1.8, ZVal: 0}), 11.25)
	if err != nil {
		panic(err)
	}

	p := Projectile{
		Position: &tuple.Point{XVal: 0, YVal: 1, ZVal: 0},
		Velocity: velocity,
	}

	env := Environment{
		Gravity: tuple.Vector{XVal: 0, YVal: -0.1, ZVal: 0},
		Wind:    tuple.Vector{XVal: -0.01, YVal: 0, ZVal: 0},
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
