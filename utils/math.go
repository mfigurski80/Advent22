package utils

import "golang.org/x/exp/constraints"

func Min[A constraints.Ordered](a, b A) A {
	if a < b {
		return a
	}
	return b
}

func Max[A constraints.Ordered](a, b A) A {
	if a > b {
		return a
	}
	return b
}

func Abs[A int | uint](a A) A {
	if a < 0 {
		return -a
	}
	return a
}
