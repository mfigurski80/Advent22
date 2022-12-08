package d4

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/mfigurski80/AOC22/utils"
)

func parseUint(s string) uint {
	v, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		panic(err)
	}
	return uint(v)
}

func parseRanges(line string) ([2]uint, [2]uint) {
	assignments := strings.Split(line, ",")
	range_a := strings.Split(assignments[0], "-")
	range_b := strings.Split(assignments[1], "-")
	return [2]uint{parseUint(range_a[0]), parseUint(range_a[1])}, [2]uint{parseUint(range_b[0]), parseUint(range_b[1])}
}

func Main() {
	n_overlap := 0
	DoByFileLine("d4/in.txt", func(line string) {
		range_a, range_b := parseRanges(line)
		if (range_a[0] <= range_b[0] && range_b[0] <= range_a[1]) || (range_a[0] <= range_b[1] && range_b[1] <= range_a[1]) ||
			(range_b[0] <= range_a[0] && range_a[0] <= range_b[1]) || (range_b[0] <= range_a[1] && range_a[1] <= range_b[1]) {
			fmt.Printf("OVERLAPPING: %d-%d, %d-%d\n", range_a[0], range_a[1], range_b[0], range_b[1])
			n_overlap++
		} else {
			fmt.Printf("NON-OVERLAPPING: %d-%d, %d-%d\n", range_a[0], range_a[1], range_b[0], range_b[1])
		}
	})
	fmt.Printf("N OVERLAPPING: %d\n", n_overlap)
}
