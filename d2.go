package main

import (
	"fmt"
	"strings"
)

const (
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3
)

func map_play(s string) uint8 {
	switch s {
	case "A":
		fallthrough
	case "X":
		return ROCK
	case "B":
		fallthrough
	case "Y":
		return PAPER
	case "C":
		fallthrough
	case "Z":
		return SCISSORS
	default:
		panic(fmt.Sprintf("Invalid play {%s}", s))
	}
}

func d2() {

	score := 0
	do_by_file_line("input/d2.txt", func(s string) {
		sp := strings.Split(s, " ")

		you := map_play(sp[0])
		me := map_play(sp[1])

		var result uint8 = 0
		if me == you {
			result = 3
		} else if me == you+1 || (me == 1 && you == 3) {
			result = 6
		}
		fmt.Printf("%d : %d -> %d (%d)\n", you, me, result, me+result)
		score += int(me + result)
	})
	fmt.Println(score)
}
