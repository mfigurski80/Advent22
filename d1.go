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

func main() {
	fmt.Println("vim-go")
	// read d1data.txt
	f, err := os.Open("d1data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	max := 0
	current := 0
	do_by_line(f, func(s string) {
		// convert to int
		i, err := strconv.Atoi(s)
		if err != nil {
			// must be ""
			if current > max {
				max = current
			}
			current = 0
		}
		current += i
	})

	fmt.Println(max)
}
