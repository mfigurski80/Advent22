package utils

import "fmt"

type QueueStruct[T any] struct {
	Push   func(T)
	Pop    func() T
	Peek   func() *T
	Length func() int
	String func() string
	Array  *[]T
}

func Queue[T any]() QueueStruct[T] {
	slice := make([]T, 0)
	return QueueStruct[T]{
		Push: func(i T) {
			slice = append(slice, i)
		},
		Pop: func() T {
			res := slice[0]
			slice = slice[1:]
			return res
		},
		Peek: func() *T {
			return &slice[0]
		},
		Length: func() int {
			return len(slice)
		},
		String: func() string {
			return fmt.Sprint(slice)
		},
		Array: &slice,
	}
}
