package utils

import "fmt"

type StackStruct[T any] struct {
	Push   func(T)
	Pop    func() T
	Length func() int
	String func() string
	Array  *[]T
}

func Stack[T any]() StackStruct[T] {
	slice := make([]T, 0)
	return StackStruct[T]{
		Push: func(i T) {
			slice = append(slice, i)
		},
		Pop: func() T {
			res := slice[len(slice)-1]
			slice = slice[:len(slice)-1]
			return res
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
