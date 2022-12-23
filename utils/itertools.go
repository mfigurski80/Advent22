package utils

import "fmt"

func FindTopK(l []int, k int) []int {
	// find top k by partitioning
	pivot := l[0]
	left := make([]int, 0)
	right := make([]int, 0)
	for _, v := range l[1:] {
		if v > pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	if len(left) == k-1 {
		return append(left, pivot)
	}
	if len(left) > k-1 {
		return FindTopK(left, k)
	}
	return append(left, append([]int{pivot}, FindTopK(right, k-len(left)-1)...)...)
}

func ReverseInPlace[T any](l []T) []T {
	// in-place reverse
	for i := 0; i < len(l)/2; i++ {
		l[i], l[len(l)-1-i] = l[len(l)-1-i], l[i]
	}
	return l
}

type StopIterationError struct {
	Message string
}

func (e StopIterationError) Error() string {
	return fmt.Sprintf("Iteration stopped: %s", e.Message)
}
