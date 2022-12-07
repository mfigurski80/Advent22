package main

// https://adventofcode.com/2022/day/1

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func d1() {
	// open d1data.txt
	f, err := os.Open("input/d1.txt")
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
