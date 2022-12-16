package d12

import (
	"fmt"

	. "github.com/mfigurski80/AOC22/utils"
)

type Position struct {
	X, Y, Height uint
}

func buildGraph(fname string) (*GraphNode[Position], *GraphNode[Position], error) {
	m := [][]GraphNode[Position]{}
	var start, end *GraphNode[Position] = nil, nil
	width := 0
	DoByFileLine(fname, func(line string) {
		if width == 0 {
			width = len(line)
		}
		m = append(m, make([]GraphNode[Position], width))
		for i, c := range line {
			height := uint(c - 'a')
			m[len(m)-1][i] = GraphNode[Position]{
				To: nil, From: nil,
				Metadata: Position{uint(i), uint(len(m) - 1), height},
			}
			cur := &m[len(m)-1][i]
			if c == 'S' {
				start = cur
				start.Metadata.Height = 0
			} else if c == 'E' {
				end = cur
				end.Metadata.Height = uint('z' - 'a')
			}
			// connect to previous item
			if i > 0 {
				prev := &m[len(m)-1][i-1]
				if prev.Metadata.Height <= height+1 {
					cur.To = append(cur.To, prev)
					prev.From = append(prev.From, cur)
				}
				if height <= prev.Metadata.Height+1 {
					prev.To = append(prev.To, cur)
					cur.From = append(cur.From, prev)
				}
			}
			// connect to previous row
			if len(m) > 1 {
				prev := &m[len(m)-2][i]
				if prev.Metadata.Height <= height+1 {
					cur.To = append(cur.To, prev)
					prev.From = append(prev.From, cur)
				}
				if height <= prev.Metadata.Height+1 {
					prev.To = append(prev.To, cur)
					cur.From = append(cur.From, prev)
				}
			}
		}
	})
	return start, end, nil
}

func Main() {
	// build terrain graph
	start, end, err := buildGraph("d12/in.txt")
	if err != nil {
		panic(err)
	}
	// find shortest path to end node
	path := uint(0)
	err = start.BfsOnGraph(0, func(node *GraphNode[Position], level uint) error {
		// fmt.Println(node.Metadata, level)
		if node.Metadata.Height == end.Metadata.Height {
			path = level
			return StopIterationError{Message: "Found exit node"}
		}
		return nil
	})
	if err != nil {
		if _, ok := err.(StopIterationError); !ok {
			panic(err)
		}
	}
	fmt.Println("Distance to End:", path)
}
