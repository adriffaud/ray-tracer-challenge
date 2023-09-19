package internal

import (
	"math"
	"testing"
)

func TestCameraBuilding(t *testing.T) {
	c := NewCamera(160, 120, math.Pi/2)

	if c.Width != 160 {
		t.Fatalf("expected 160. got=%d", c.Width)
	}
	if c.Height != 120 {
		t.Fatalf("expected 120. got=%d", c.Height)
	}
	if c.FieldOfView != math.Pi/2 {
		t.Fatalf("expected Pi/2. got=%f", c.FieldOfView)
	}
	if !Eq(c.Transform, IdentityMatrix()) {
		t.Fatalf("expected IdentityMatrix. got=%+v", c.Transform)
	}
}

func TestCameraPixelSize(t *testing.T) {
	tests := []struct {
		camera   Camera
		expected float64
	}{
		{NewCamera(200, 125, math.Pi/2), 0.01},
		{NewCamera(125, 200, math.Pi/2), 0.01},
	}

	for _, test := range tests {
		actual := test.camera.PixelSize
		if !ApproxEq(actual, test.expected) {
			t.Fatalf("expected %f. got %f", test.expected, actual)
		}
	}
}

func TestRayFromCamera(t *testing.T) {
	tests := []struct {
		transform Matrix
		c         Camera
		x, y      int
		e         Ray
	}{
		{
			IdentityMatrix(),
			NewCamera(201, 101, math.Pi/2),
			100, 50,
			Ray{Origin: Point{}, Direction: Vector{Z: -1}},
		},
		{
			IdentityMatrix(),
			NewCamera(201, 101, math.Pi/2),
			0, 0,
			Ray{Origin: Point{}, Direction: Vector{X: 0.66519, Y: 0.33259, Z: -0.66851}},
		},
		{
			RotationY(math.Pi / 4).Multiply(Translation(0, -2, 5)),
			NewCamera(201, 101, math.Pi/2),
			100, 50,
			Ray{Origin: Point{Y: 2, Z: -5}, Direction: Vector{X: math.Sqrt(2) / 2, Z: -math.Sqrt(2) / 2}},
		},
	}

	for _, test := range tests {
		test.c.Transform = test.transform
		r := test.c.RayForPixel(test.x, test.y)

		assertPointEquals(t, test.e.Origin, r.Origin)
		assertVectorEquals(t, test.e.Direction, r.Direction)
	}
}

func TestRenderingWorldWithCamera(t *testing.T) {
	w := NewWorld()
	c := NewCamera(11, 11, math.Pi/2)
	c.Transform = ViewTransform(Point{Z: -5}, Point{}, Vector{Y: 1})
	image := c.Render(w)
	p := image.PixelAt(5, 5)
	expected := Color{R: 0.38066, G: 0.47583, B: 0.2855}

	assertColorEquals(t, expected, p)
}
