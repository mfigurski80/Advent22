package d3

import (
	"fmt"
	"math"

	. "github.com/mfigurski80/AOC22/utils"
)

type CharSet uint64

const FULL_CHARSET CharSet = math.MaxUint64

func getPriority(c rune) uint {
	if c >= 'a' && c <= 'z' {
		return uint(c - 'a' + 1)
	}
	if c >= 'A' && c <= 'Z' {
		return uint(c - 'A' + 27)
	}
	panic("invalid char: " + string(c))
}

func getSet(line string) CharSet {
	var set CharSet = 0
	for _, c := range line {
		set |= 1 << getPriority(c)
	}
	return set
}

func Main() {
	sum_priorities := uint(0)

	DoByNFileLines(3, "d3/in.txt", func(lines []string) {
		var set CharSet = FULL_CHARSET
		for _, line := range lines {
			set &= getSet(line)
		}
		if set == 0 {
			panic("nothing in common")
		}
		pos := uint(math.Log2(float64(set)))
		sum_priorities += pos
		fmt.Printf("set: %064b (%d)\n", set, pos)
	})

	fmt.Printf("sum: %d\n", sum_priorities)
}
