package shape

import (
	"sort"
)

type Intersection struct {
	Object Shape
	T      float64
}

type Intersections []Intersection

func (xs Intersections) Hit() Intersection {
	var positives Intersections

	for _, i := range xs {
		if i.T >= 0 {
			positives = append(positives, i)
		}
	}

	sort.Slice(positives, func(i, j int) bool {
		return positives[i].T < positives[j].T
	})

	if len(positives) > 0 {
		return positives[0]
	}
	return Intersection{}
}
