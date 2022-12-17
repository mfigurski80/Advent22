package d14

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/mfigurski80/AOC22/utils"
)

type Point [2]uint // x, y
type Path = []Point

type Tile uint8
type World [][]Tile
type ScaleFunc func(uint) uint

const (
	EMPTY_TILE Tile = ' '
	ROCK_TILE  Tile = '#'
	SAND_TILE  Tile = '.'
)

func (w *World) String() string {
	str := fmt.Sprintf("World %d x %d:\n", len(*w), len((*w)[0]))
	for i, row := range *w {
		str += fmt.Sprintf(" %-2d|%s|\n", i, string(row))
	}
	str += "   +"
	for i := 0; i < len((*w)[0]); i += 2 {
		str += fmt.Sprintf("%d-", i%10)
	}
	str += "+"

	return str
}

func ReadScaleWorldFrom(fname string) (*World, ScaleFunc) {
	rockPaths := make([]Path, 0)
	DoByFileLine("d14/in.txt", func(line string) {
		points := strings.Split(line, " -> ")
		path := make(Path, len(points))
		for i, p := range points {
			coords := strings.Split(p, ",")
			x, err := strconv.Atoi(coords[0])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(coords[1])
			if err != nil {
				panic(err)
			}
			path[i] = Point{uint(x), uint(y)}
		}
		rockPaths = append(rockPaths, path)
	})

	// create world of optimal size
	maxY := uint(0)
	for _, path := range rockPaths {
		for _, p := range path {
			if p[1] > maxY {
				maxY = p[1]
			}
		}
	}
	xScale := ScaleFunc(func(x uint) uint {
		return x - 500 + maxY
	})
	world := make(World, maxY+1)
	for y := range world {
		world[y] = make([]Tile, 2*maxY)
		for x := range world[y] {
			world[y][x] = EMPTY_TILE
		}
	}

	// fill world with rocks
	for _, path := range rockPaths {
		// interpolate between points
		for i := 0; i < len(path)-1; i++ {
			p1 := path[i]
			p2 := path[i+1]
			// fmt.Printf("Interpolating between %v and %v\n", path[i], path[i+1])
			if p1[0] == p2[0] {
				// vertical line
				st := Min(p1[1], p2[1])
				end := Max(p1[1], p2[1])
				for y := st; y <= end; y++ {
					world[y][xScale(p1[0])] = ROCK_TILE
				}
			} else if p1[1] == p2[1] {
				// horizontal line
				st := Min(p1[0], p2[0])
				end := Max(p1[0], p2[0])
				for x := st; x <= end; x++ {
					world[p1[1]][xScale(x)] = ROCK_TILE
				}
			} else {
				// diagonal line
				panic("diagonal line")
			}
		}
		// for _, p := range path {
		// 	world[p[1]][xScale(p[0])] = ROCK_TILE
		// }
	}

	return &world, xScale
}

func Main() {
	world, xScale := ReadScaleWorldFrom("d14/in.txt")
	// (*world)[8][9] = SAND_TILE

	nPlaced := uint(0)
	path := Stack[Point]()
	path.Push(Point{xScale(500), 0})
	for path.Length() > 0 {
		p := path.Peek()
		// if down goes off map, done
		if p[1]+1 >= uint(len(*world)) {
			break
		}
		// if down is empty, go down
		if (*world)[p[1]+1][p[0]] == EMPTY_TILE {
			path.Push(Point{p[0], p[1] + 1})
			continue
		}
		// if down left goes off map, done
		if p[0] == 0 {
			break
		}
		// if down left is empty, go down left
		if (*world)[p[1]+1][p[0]-1] == EMPTY_TILE {
			path.Push(Point{p[0] - 1, p[1] + 1})
			continue
		}
		// if down right goes off map, done
		if p[0] == uint(len((*world)[0]))-1 {
			break
		}
		// if down right is empty, go down right
		if (*world)[p[1]+1][p[0]+1] == EMPTY_TILE {
			path.Push(Point{p[0] + 1, p[1] + 1})
			continue
		}
		// if can't move, drop sand and pop
		(*world)[p[1]][p[0]] = SAND_TILE
		nPlaced++
		path.Pop()
	}

	fmt.Println(world)
	fmt.Println(nPlaced)
}
