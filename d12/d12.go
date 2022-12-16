package d12

import (
	"fmt"
	"math"

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
				height = 0
			} else if c == 'E' {
				end = cur
				end.Metadata.Height = uint('z' - 'a')
				height = uint('z' - 'a')
			}
			// connect to previous item
			if i > 0 {
				prev := &m[len(m)-1][i-1]
				if prev.Metadata.Height <= height+1 { // if previous tile is reachable (increase at most 1)
					cur.To = append(cur.To, prev)      // current tile goes to prev
					prev.From = append(prev.From, cur) // prev comes from current
				} // note this can be true as well!
				if height <= prev.Metadata.Height+1 { // if previous tile can reach (increase at most 1)
					prev.To = append(prev.To, cur)    // prev tile goes to current
					cur.From = append(cur.From, prev) // current comes from prev
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
	// find shortest path to end node, starting from max uint
	var shortest_path uint = math.MaxUint
	BfsCore(end, func(node *GraphNode[Position], level uint) ([]*GraphNode[Position], error) {
		if node.Metadata.Height == start.Metadata.Height && level < shortest_path {
			fmt.Printf(" -- Found path of length %d at (%d, %d)\n", level, node.Metadata.X, node.Metadata.Y)
			shortest_path = level
		}
		children := []*GraphNode[Position]{}
		for i := range node.From {
			if node.From[i].VisitedDepth > 0 {
				continue
			}
			node.From[i].VisitedDepth = level + 1
			children = append(children, node.From[i])
		}
		return children, nil
	})
	// find shortest-path node with height of start
	fmt.Printf("Shortest path: %d\n", shortest_path)
	// fmt.Println("From", end.From, ", To", end.To, ", Height", end.Metadata.Height)
	// left := *(end.To[0])
	// fmt.Println("Left From", left.From, ", To", left.To, ", Height", left.Metadata.Height)
}
