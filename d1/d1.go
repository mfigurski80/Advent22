package d1

// https://adventofcode.com/2022/day/1

import (
	"fmt"
	"log"
	"os"
	"strconv"

	. "github.com/mfigurski80/AOC22/utils"
)

func Main() {
	// open d1data.txt
	f, err := os.Open("d1/d1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// read to get calories being carried
	carried := make([]int, 0)
	current := 0
	DoByLine(f, func(s string) {
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
	top3 := FindTopK(carried, 3)
	sum := top3[0] + top3[1] + top3[2]
	fmt.Println(sum)
}
