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
	// get range covered by sensor at y
	offset := d - dy
	if offset < 0 {
		return false, Range[int]{}
	}
	return true, Range[int]{Lo: s.Pos.X - offset, Hi: s.Pos.X + offset}
}

func tuningFrequency(x, y int) int {
	return x*4000000 + y
}

var TARGET_RANGE = Range[int]{Lo: 0, Hi: 4000000}
var TARGET_AREA = TARGET_RANGE.Area()

func Main() {
	// read sensors
	sensors := make([]Sensor, 0)
	forSensorInFile("d15/in.txt", func(s Sensor) {
		sensors = append(sensors, s)
	})

	// for each possible y
	for y := TARGET_RANGE.Lo; y < TARGET_RANGE.Hi; y++ {
		// build ranges
		ranges := make([]Range[int], len(sensors))
		for i, s := range sensors {
			ok, r := getRangeCoveredAtY(y, s)
			if !ok {
				continue
			}
			ranges[i] = r
		}
		// combine ranges
		combined := CombineRangeSeries(ranges)
		// check area
		area := 0
		for _, r := range combined {
			area += r.And(TARGET_RANGE).Area()
		}
		if y%10000 == 0 {
			fmt.Printf("%-2d Total range covered: %d / %d\n", y, area, TARGET_AREA)
		}
		if area == TARGET_AREA {
			continue
		}
		fmt.Printf("FOUND Not covering full range at y=%d, x=%d?\n",
			y, combined[0].Hi+1)
		fmt.Println(combined)
		fmt.Printf("Tuning Frequency of this coordinate: %d\n", tuningFrequency(combined[0].Hi+1, y))
		break
	}

}
