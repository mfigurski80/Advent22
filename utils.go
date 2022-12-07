package main

import (
	"bufio"
	"os"
)

func do_by_line(f *os.File, fn func(string)) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fn(scanner.Text())
	}
}

func do_by_file_line(f string, fn func(string)) {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	do_by_line(file, fn)
}

func find_top_k(l []int, k int) []int {
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
		return find_top_k(left, k)
	}
	return append(left, append([]int{pivot}, find_top_k(right, k-len(left)-1)...)...)
}
