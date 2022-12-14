package d10

import (
	"fmt"
	"regexp"
	"strconv"

	. "github.com/mfigurski80/AOC22/utils"
)

var addInstr = regexp.MustCompile(`addx (?P<Val>-?\d+)`)

var REPORT = []uint{20, 60, 100, 140, 180, 220, 100000000000}

func printFromCycle(cycle int, x int) {
	cycle = cycle % 40
	if cycle == x || cycle == x+1 || cycle == x-1 {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
	if cycle == 39 {
		fmt.Println("")
	}
}

func Main() {
	x := 1
	cycle := uint(0)
	DoByFileLine("d10/in.txt", func(line string) {
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
			printFromCycle(int(cycle), x)
			printFromCycle(int(cycle+1), x)
			cycle += 2
			x += val
		case line == "noop":
			printFromCycle(int(cycle), x)
			cycle++
		default:
			fmt.Println("unknown instruction", line)
			panic("unknown instruction")
		}
	})

}
