package main

// https://adventofcode.com/2022/day/1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func do_by_line(f *os.File, fn func(string)) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fn(scanner.Text())
	}
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

func d1() {
	fmt.Println("vim-go")
	// open d1data.txt
	f, err := os.Open("d1data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// read to get calories being carried
	carried := make([]int, 0)
	current := 0
	do_by_line(f, func(s string) {
		// convert to int
		i, err := strconv.Atoi(s)
		if err != nil {
			// must be ""
			carried = append(carried, current)
			current = 0
		} else {
			current += i
		}
	})
	if current != 0 {
		carried = append(carried, current)
	}

	// find sum of top 3
	top3 := find_top_k(carried, 3)
	sum := top3[0] + top3[1] + top3[2]
	fmt.Println(sum)
}
