package shape

import (
	"sort"

	"github.com/adriffaud/ray-tracer-challenge/pkg/primitives"
)

type Computations struct {
	Intersection
	Inside  bool
	Point   primitives.Point
	EyeV    primitives.Vector
	NormalV primitives.Vector
}

type Intersection struct {
	Object   Shape
	Distance float64
}

type Intersections []Intersection

func (i Intersection) PrepareComputations(r primitives.Ray) Computations {
	p := r.Position(i.Distance)

	comps := Computations{
		Intersection: i,
		Point:        p,
		EyeV:         r.Direction.Negate(),
		NormalV:      Sphere().NormalAt(p),
	}

	if comps.NormalV.Dot(comps.EyeV) < 0 {
		comps.Inside = true
		comps.NormalV = comps.NormalV.Negate()
	}

	return comps
}

func (i Intersections) Len() int { return len(i) }

func (i Intersections) Less(j, k int) bool { return i[j].Distance < i[k].Distance }

func (i Intersections) Swap(j, k int) { i[j], i[k] = i[k], i[j] }

func (xs Intersections) Hit() (Intersection, bool) {
	sort.Sort(xs)
	for _, i := range xs {
		if i.Distance >= 0 {
			return i, true
		}
	}

	return Intersection{}, false
}