package d2

import (
	"fmt"
	"strings"

	. "github.com/mfigurski80/AOC22/utils"
)

// moves/results type
const (
	ROCK = iota
	PAPER
	SCISSORS
)
const (
	LOSE = iota
	DRAW
	WIN
)

// scoring vectors
var MOVE_SCORE = [3]int8{1, 2, 3}
var RESULT_SCORE = [3]int8{0, 3, 6}

func map_move(s string) int8 {
	switch s {
	case "A":
		return ROCK
	case "B":
		return PAPER
	case "C":
		return SCISSORS
	default:
		panic(fmt.Sprintf("Invalid move {%s}", s))
	}
}

func map_result(s string) int8 {
	switch s {
	case "X":
		return LOSE
	case "Y":
		return DRAW
	case "Z":
		return WIN
	default:
		panic(fmt.Sprintf("Invalid result {%s}", s))
	}
}

func parse_move_for_result(a_move int8, result int8) int8 {
	var b_move int8 = 0
	if result == LOSE {
		// a_move = 1, lose so play 2
		// a_move = 2, play 3
		// a_move = 3, play 1
		b_move = (a_move + 2) % 3
	} else if result == DRAW {
		b_move = a_move
	} else {
		// a_move = 1, we need to win, so play 3
		// a_move = 2, play 1
		// a_move = 3, play 2
		b_move = (a_move + 1) % 3
	}
	return b_move
}

func Main() {
	score := 0
	DoByFileLine("d2/d2.txt", func(s string) {
		sp := strings.Split(s, " ")
		a_move := map_move(sp[0])
		result := map_result(sp[1])
		b_move := parse_move_for_result(a_move, result)
		fmt.Printf("%d, %d: %d (%d)\n", a_move, b_move, result, RESULT_SCORE[result]+MOVE_SCORE[b_move])
		score += int(RESULT_SCORE[result] + MOVE_SCORE[b_move])
	})
	fmt.Println(score)
}
