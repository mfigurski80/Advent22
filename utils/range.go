package utils

import (
	"fmt"
	"sort"
)

// make type for int | uint
type Number interface {
	int | uint
}

type Range[T Number] struct {
	Lo, Hi T
}

func (r Range[T]) String() string {
	return fmt.Sprintf("[%v,%v]", r.Lo, r.Hi)
}

func (r Range[T]) Or(other Range[T]) Range[T] {
	if r.Lo > other.Hi+1 || r.Hi+1 < other.Lo {
		return r
	}
	return Range[T]{Lo: Min(r.Lo, other.Lo), Hi: Max(r.Hi, other.Hi)}
}

func (r Range[T]) Contiguous(other Range[T]) bool {
	return !(r.Lo > other.Hi+1 || r.Hi+1 < other.Lo)
}

func CombineRangeSeries[T Number](s []Range[T]) []Range[T] {
	// sort ranges by lo
	sort.Slice(s, func(i, j int) bool {
		return s[i].Lo < s[j].Lo
	})
	// combine ranges, while they're still converging
	prevLength := len(s)
	var combined = []Range[T]{}
	for len(combined) < prevLength {
		combined = []Range[T]{s[0]}
		for _, r := range s[1:] { // for each range to combine
			for i, c := range combined { // for each already-combined range
				if c.Contiguous(r) { // if overlapping, combine
					combined[i] = c.Or(r)
					fmt.Printf("Combined ranges %v and %v to %v\n", c, r, combined[i])
					goto next // found, stop looking for combined ranges
				} else {
					// fmt.Printf("Ranges %v and %v do not overlap\n", c, r)
				}
			}
			combined = append(combined, r)
		next:
		}
		prevLength = len(s)
		s = combined
		if len(s) == 1 {
			break
		}
	}
	return combined
}
