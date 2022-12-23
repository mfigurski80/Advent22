package utils

import "golang.org/x/exp/constraints"

/**
 * Heap data structure
 * @param T: type of the data stored in the heap
 */
type HeapStruct[T comparable] struct {
	// NOTE: this is a min-heap
	// NOTE: this is a binary heap
	Array  []T
	Push   func(T) bool
	Pop    func() T
	Peek   func() *T
	Length func() uint
}

func Heap[T constraints.Ordered]() HeapStruct[T] {
	slice := make([]T, 0)
	return HeapStruct[T]{
		Array: slice,
		Push: func(v T) bool {
			// add to end
			slice = append(slice, v)
			// bubble up
			i := len(slice) - 1
			for i > 0 {
				parent := (i - 1) / 2
				if slice[parent] < slice[i] {
					return false
				}
				slice[parent], slice[i] = slice[i], slice[parent]
				i = parent
			}
			return true
		},
		Pop: func() T {
			// pop root
			root := slice[0]
			// move last to root
			slice[0] = slice[len(slice)-1]
			slice = slice[:len(slice)-1]
			// bubble down
			i := 0
			for i < len(slice) {
				left := 2*i + 1
				right := 2*i + 2
				if left >= len(slice) {
					break
				}
				if right >= len(slice) {
					if slice[left] < slice[i] {
						slice[left], slice[i] = slice[i], slice[left]
					}
					break
				}
				if slice[left] < slice[i] || slice[right] < slice[i] {
					if slice[left] < slice[right] {
						slice[left], slice[i] = slice[i], slice[left]
						i = left
					} else {
						slice[right], slice[i] = slice[i], slice[right]
						i = right
					}
				} else {
					break
				}
			}
			return root
		},
		Peek: func() *T {
			return &slice[0]
		},
		Length: func() uint {
			return uint(len(slice))
		},
	}
}

/**
 * Function to build a heap from an array
 * @param T: type of the data stored in the heap
 * @param arr: array to build the heap from
 */
func Heapify[T constraints.Ordered](slice []T) HeapStruct[T] {
	heap := Heap[T]()
	for _, v := range slice {
		heap.Push(v)
	}
	return heap
}

/**
 * Utility to enable heap with non-comparable types
 * @param T: type of the data stored in the heap
 * @param V: type of the value used to score the data
 */
type HeapValueWrapper[T any, V constraints.Ordered] struct {
	Element T
	Value   V
	Depth   uint
}

func (h *HeapValueWrapper[T, V]) Less(other *HeapValueWrapper[T, V]) bool {
	return h.Value < other.Value
}
