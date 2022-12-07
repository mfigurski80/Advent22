package d3

import (
	"fmt"

	. "github.com/mfigurski80/AOC22/utils"
)

func getPriority(c rune) uint {
	if c >= 'a' && c <= 'z' {
		return uint(c - 'a' + 1)
	}
	if c >= 'A' && c <= 'Z' {
		return uint(c - 'A' + 27)
	}
	panic("invalid char: " + string(c))
}

func Main() {
	sum_priorities := uint(0)

	DoByFileLine("d3/in.txt", func(line string) {
		if len(line)%2 != 0 {
			panic("line length must be even: " + line)
		}
		// build first-half set for chars
		// use uint64 and bit-shift to build set
		cmpt_a := uint64(0)
		for _, c := range line[:len(line)/2] {
			cmpt_a |= 1 << getPriority(c)
		}
		// fmt.Printf("cmpt_a: %064b\n", cmpt_a)
		// check for duplicates in second-half
		for _, c := range line[len(line)/2:] {
			if cmpt_a&(1<<getPriority(c)) != 0 {
				sum_priorities += getPriority(c)
				fmt.Printf("duplicate: %c, %d\n", c, getPriority(c))
				break
			}
		}
	})
	fmt.Printf("sum: %d\n", sum_priorities)
}
