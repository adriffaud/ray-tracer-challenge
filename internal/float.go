package internal

import "math"

const EPSILON = 1.0e-4

// ApproxEq compares two float64 and determine if they are close enough to be considered equals.
// Inspiration : https://floating-point-gui.de/errors/comparison/
func ApproxEq(a, b float64) bool {
	absA := math.Abs(a)
	absB := math.Abs(b)
	diff := math.Abs(a - b)

	if a == b {
		return true
	} else if a == 0 || b == 0 {
		return diff < EPSILON
	} else {
		return diff/math.Min(absA+absB, math.MaxFloat64) < EPSILON
	}
}
