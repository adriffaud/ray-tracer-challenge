package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"

	"github.com/adriffaud/ray-tracer-challenge/internal"
)

type XYRay struct {
	X, Y int
}

const (
	WIDTH  = 1024
	HEIGHT = 768
)

var (
	imgMutex *sync.Mutex
	ch       chan (XYRay)
	wg       sync.WaitGroup
	world    internal.World
	camera   internal.Camera
	canvas   internal.Canvas
)

func RenderPixel() {
	defer wg.Done()
	open := true
	xyray := XYRay{}
	for open {
		xyray, open = <-ch
		ray := camera.RayForPixel(xyray.X, xyray.Y)
		col := world.ColorAt(ray)

		imgMutex.Lock()
		canvas.WritePixel(xyray.X, xyray.Y, col)
		imgMutex.Unlock()
	}
}

func main() {
	fmt.Println("Starting render")

	ch = make(chan XYRay, 1000)
	imgMutex = &sync.Mutex{}
	start := time.Now()

	canvas = *internal.NewCanvas(WIDTH, HEIGHT)

	camera = internal.NewCamera(WIDTH, HEIGHT, math.Pi/3)
	camera.Transform = internal.ViewTransform(internal.Point{Y: 1.5, Z: -5}, internal.Point{Y: 1}, internal.Vector{Y: 1})

	floorMaterial := internal.NewMaterial()
	floorMaterial.Color = internal.Color{R: 1, G: 0.9, B: 0.9}
	floorMaterial.Specular = 0

	floor := internal.Sphere()
	floor.Transform = internal.Scaling(10, 0.01, 10)
	floor.Material = floorMaterial

	m1 := internal.Translation(0, 0, 5).Multiply(internal.RotationY(-math.Pi / 4))
	m2 := m1.Multiply(internal.RotationX(math.Pi / 2))
	m3 := m2.Multiply(internal.Scaling(10, 0.01, 10))
	leftWall := internal.Sphere()
	leftWall.Transform = m3
	leftWall.Material = floorMaterial

	rightWall := internal.Sphere()
	rightWall.Transform = internal.Translation(0, 0, 5).Multiply(
		internal.RotationY(math.Pi / 4).Multiply(
			internal.RotationX(math.Pi / 2)).Multiply(
			internal.Scaling(10, 0.01, 10)))
	rightWall.Material = floorMaterial

	middle := internal.Sphere()
	middle.Transform = internal.Translation(-0.5, 1, 0.5)
	middle.Material.Color = internal.Color{R: 0.1, G: 1, B: 0.5}
	middle.Material.Diffuse = 0.7
	middle.Material.Specular = 0.3

	right := internal.Sphere()
	right.Transform = internal.Translation(1.5, 0.5, -0.5).Multiply(internal.Scaling(0.5, 0.5, 0.5))
	right.Material.Color = internal.Color{R: 0.5, G: 1, B: 0.1}
	right.Material.Diffuse = 0.7
	right.Material.Specular = 0.3

	left := internal.Sphere()
	left.Transform = internal.Translation(-1.5, 0.33, -0.75).Multiply(internal.Scaling(0.33, 0.33, 0.33))
	left.Material.Color = internal.Color{R: 1, G: 0.8, B: 0.1}
	left.Material.Diffuse = 0.7
	left.Material.Specular = 0.3

	world = internal.World{
		Lights: []internal.Light{{
			Position:  internal.Point{X: -10, Y: 10, Z: -10},
			Intensity: internal.Color{R: 1, G: 1, B: 1},
		}},
		Objects: []internal.Shape{floor, leftWall, rightWall, middle, right, left},
	}

	cpus := runtime.NumCPU()
	wg.Add(cpus)
	for t := 0; t < cpus; t++ {
		go RenderPixel()
	}
	fmt.Println("Starting pixel calculations")
	for y := 0; y < camera.Height; y++ {
		for x := 0; x < camera.Width; x++ {
			ch <- XYRay{X: x, Y: y}
		}
	}
	fmt.Println("Closing channel")
	close(ch)
	wg.Wait()

	fmt.Printf("Render finished: %v\n", time.Since(start))
	internal.Export(&canvas, "world.png")
}
