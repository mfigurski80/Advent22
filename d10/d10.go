package d10

import (
	"fmt"
	"regexp"
	"strconv"

	. "github.com/mfigurski80/AOC22/utils"
)

var addInstr = regexp.MustCompile(`addx (?P<Val>-?\d+)`)

var REPORT = []uint{20, 60, 100, 140, 180, 220, 100000000000}

func Main() {
	x := 1
	cycle := uint(0)
	nextReport := 0
	sumStrength := 0
	DoByFileLine("d9/in.txt", func(line string) {
		switch {
		case addInstr.MatchString(line):
			match := addInstr.FindStringSubmatch(line)
			if len(match) != 2 {
				fmt.Println(match)
				panic("bad match")
			}
			val, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			cycle += 2
			if REPORT[nextReport] <= cycle {
				fmt.Printf("[REPORT %d]: x = %d (strength %d)\n", REPORT[nextReport], x, x*int(REPORT[nextReport]))
				sumStrength += x * int(REPORT[nextReport])
				nextReport++
			}
			x += val
		case line == "noop":
			cycle++
			if REPORT[nextReport] <= cycle {
				fmt.Printf("[REPORT %d]: x = %d (strength %d)\n", REPORT[nextReport], x, x*int(REPORT[nextReport]))
				sumStrength += x * int(REPORT[nextReport])
				nextReport++
			}
		default:
			fmt.Println("unknown instruction", line)
			panic("unknown instruction")
		}
	})

	fmt.Println("Total Strength", sumStrength)
}
