package d6

import (
	"fmt"

	"math/bits"

	. "github.com/mfigurski80/AOC22/utils"
)

type CharMap uint32

func LetterToMap(c rune) CharMap {
	return CharMap(1 << (c - 'a'))
}
func StringToMap(s string) CharMap {
	var res CharMap
	for _, c := range s {
		res |= LetterToMap(c)
	}
	return res
}

var LOOK_FOR_COUNT = 14

func Main() {
	DoByFileLine("d6/in.txt", func(ln string) {
		for i := LOOK_FOR_COUNT; i < len(ln); i++ {
			m := StringToMap(ln[i-LOOK_FOR_COUNT : i])
			// github.com/tmthrgd/go-popcount
			// https://pkg.go.dev/math/bits#OnesCount32
			count := bits.OnesCount32(uint32(m))
			// fmt.Printf("Checking %s at %02d: %032b (%d)\n", ln[i-LOOK_FOR_COUNT:i], i, m, count)
			if count == LOOK_FOR_COUNT {
				fmt.Printf("Found: %s at %d: %b\n", ln[i-LOOK_FOR_COUNT:i], i, m)
				return
			}
		}

	})
}
