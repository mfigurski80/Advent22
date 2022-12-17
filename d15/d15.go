package d15

import (
	"fmt"
	"strconv"

	. "github.com/mfigurski80/AOC22/utils"
)

type Point struct {
	X, Y int
}

type Sensor struct {
	Pos           Point
	ClosestBeacon Point
}

func forSensorInFile(fname string, f func(Sensor)) {
	DoForRegexMatches(fname, `x=(-?\d+), y=(-?\d+):.*x=(-?\d+), y=(-?\d+)`, func(matches []string) {
		a, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		c, err := strconv.Atoi(matches[3])
		if err != nil {
			panic(err)
		}
		d, err := strconv.Atoi(matches[4])
		if err != nil {
			panic(err)
		}
		f(Sensor{Pos: Point{X: a, Y: b}, ClosestBeacon: Point{X: c, Y: d}})
	})
}

func getRangeCoveredAtY(y int, s Sensor) (bool, Range[int]) {
	// get manhattan distance covered by sensor -- this is the range it can cover
	d := Abs(s.Pos.X-s.ClosestBeacon.X) + Abs(s.Pos.Y-s.ClosestBeacon.Y)
	// get distance from sensor to y -- this will decide how much range is reduced
	dy := Abs(y - s.Pos.Y)
	fmt.Println("d", d, "dy", dy)
	// get range covered by sensor at y
	offset := d - dy
	if offset < 0 {
		return false, Range[int]{}
	}
	return true, Range[int]{Lo: s.Pos.X - offset, Hi: s.Pos.X + offset}
}

const Y_TO_CHECK = 2000000

func Main() {
	// read ranges
	ranges := []Range[int]{}
	forSensorInFile("d15/in.txt", func(s Sensor) {
		ok, r := getRangeCoveredAtY(Y_TO_CHECK, s)
		if !ok {
			return
		}
		ranges = append(ranges, r)
		fmt.Printf("Sensor at %v (beacon at %v) covers range %v\n", s.Pos, s.ClosestBeacon, r)
	})

	// combine ranges
	combined := CombineRangeSeries(ranges)
	fmt.Printf("Combined ranges: %v\n", combined)
	area := 0
	for _, r := range combined {
		area += r.Hi - r.Lo
	}
	fmt.Printf("Total range covered: %d\n", area)
}
