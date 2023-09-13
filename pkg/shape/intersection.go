package shape

import (
	"sort"
)

type Intersection struct {
	Object   Shape
	Distance float64
}

type Intersections []Intersection

func (i Intersections) Len() int { return len(i) }

func (i Intersections) Less(j, k int) bool { return i[j].Distance < i[k].Distance }

func (i Intersections) Swap(j, k int) { i[j], i[k] = i[k], i[j] }

func (xs Intersections) Hit() Intersection {
	var positives Intersections

	for _, i := range xs {
		if i.Distance >= 0 {
			positives = append(positives, i)
		}
	}

	sort.Slice(positives, func(i, j int) bool {
		return positives[i].Distance < positives[j].Distance
	})

	if len(positives) > 0 {
		return positives[0]
	}
	return Intersection{}
}
