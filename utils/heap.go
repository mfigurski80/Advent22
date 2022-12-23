package utils

import "golang.org/x/exp/constraints"

/**
 * Heap data structure
 * @param T: type of the data stored in the heap
 */
type HeapStruct[T any, V constraints.Ordered] struct {
	// NOTE: this is a min-heap
	// NOTE: this is a binary heap
	Array  []*T
	Values []V
	Push   func(*T, V) bool
	Pop    func() (*T, V)
	Peek   func() (*T, V)
	Length func() uint
}

func Heap[T any, V constraints.Ordered]() HeapStruct[T, V] {
	elems := make([]*T, 0)
	values := make([]V, 0)
	return HeapStruct[T, V]{
		Array:  elems,
		Values: values,
		Push: func(e *T, v V) bool {
			// add to end
			elems = append(elems, e)
			values = append(values, v)
			// bubble up
			i := len(elems) - 1
			for i > 0 {
				parent := (i - 1) / 2
				if values[parent] < values[i] {
					return false
				}
				elems[parent], elems[i] = elems[i], elems[parent]
				values[parent], values[i] = values[i], values[parent]
				i = parent
			}
			return true
		},
		Pop: func() (*T, V) {
			// pop root
			root := elems[0]
			rootV := values[0]
			// move last to root
			last := len(elems) - 1
			elems[0] = elems[last]
			elems = elems[:last]
			values[0] = values[last]
			values = values[:last]
			// bubble down
			for i := 0; i < len(elems); {
				left := 2*i + 1
				right := 2*i + 2
				if left > len(elems) {
					break
				}
				if right >= len(elems) {
					if values[left] < values[i] {
						elems[left], elems[i] = elems[i], elems[left]
						values[left], values[i] = values[i], values[left]
					}
					break
				}
				if values[left] < values[i] || values[right] < values[i] {
					if values[left] < values[right] {
						elems[left], elems[i] = elems[i], elems[left]
						values[left], values[i] = values[i], values[left]
						i = left
					} else {
						elems[right], elems[i] = elems[i], elems[right]
						values[right], values[i] = values[i], values[right]
						i = right
					}
				} else {
					break
				}
			}
			return root, rootV
		},
		Peek: func() (*T, V) {
			return elems[0], values[0]
		},
		Length: func() uint {
			return uint(len(elems))
		},
	}
}

/**
 * Function to build a heap from an array
 * @param T: type of the data stored in the heap
 * @param arr: array to build the heap from
 */
func Heapify[T any, V constraints.Ordered](slice []*T, values []V) HeapStruct[T, V] {
	heap := Heap[T, V]()
	for i := range slice {
		heap.Push(slice[i], values[i])
	}
	return heap
}
