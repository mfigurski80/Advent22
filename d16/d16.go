package d16

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/mfigurski80/AOC22/utils"
)

type TunnelId = [2]rune

type Tunnel struct {
	Id       TunnelId
	FlowRate uint
}

const MOVE_COST = 1
const OPEN_COST = 1
const MAX_COST = 30

func buildTunnelGraphFromFile(fname string) []GraphNode[Tunnel] {
	var nodes []GraphNode[Tunnel]
	nodeIdx := make(map[TunnelId]uint)
	DoForRegexMatchesWithSetup(fname, `Valve (..) has flow rate=(\d+); tunnels? leads? to valves? (.*)`,
		func(tot uint) { // setup function
			// NOTE: don't use expandable vec because pointers don't get updated
			nodes = make([]GraphNode[Tunnel], tot)
		},
		func(m []string, nI uint) { // parsing function
			// build node
			Id := TunnelId{rune(m[1][0]), rune(m[1][1])}
			FlowRate, err := strconv.Atoi(m[2])
			if err != nil {
				panic(err)
			}
			nodes[nI] = GraphNode[Tunnel]{
				Metadata: Tunnel{
					Id:       Id,
					FlowRate: uint(FlowRate),
				},
			}
			nodeIdx[Id] = nI
			// connect edges
			for _, dest := range strings.Split(m[3], ", ") {
				destId := TunnelId{rune(dest[0]), rune(dest[1])}
				if dI, ok := nodeIdx[destId]; ok {
					nodes[nI].To = append(nodes[nI].To, &nodes[dI])
					nodes[dI].To = append(nodes[dI].To, &nodes[nI])
				}
				continue // skip destinations that don't exist yet
			}
		})
	return nodes
}

func printNode(n *GraphNode[Tunnel]) {
	fmt.Printf("  <Node: %c%c [%p] (edges ", n.Metadata.Id[0], n.Metadata.Id[1], n)
	for _, e := range n.To {
		fmt.Printf("%c%c ", e.Metadata.Id[0], e.Metadata.Id[1])
	}
	fmt.Printf(")>\n")
}

// func valueNode(n *GraphNode[Tunnel]) uint {
// val := uint(0)
// n.BfsOnGraph(0, func(n *GraphNode[Tunnel], depth uint) error {
// // printNode(n)
// if n.Metadata.Open { // no cost if already open
// return nil
// }
// val += n.Metadata.FlowRate * (MOVE_COST*depth + OPEN_COST)
// return nil
// })
// return val
// }

type State struct {
	Position         *GraphNode[Tunnel]
	ValvesOpen       []*GraphNode[Tunnel]
	PressureReleased uint
}

func Main() {
	// Build Graph
	graph := buildTunnelGraphFromFile("d16/in.txt")
	if len(graph) > 256 {
		panic("Misjudged graph size")
	}

	// Setup State Representation
	var maxState *State
	maxPressureReleased := uint(0)
	seen := map[TunnelId][]uint{} // cache: tunnel + depth => pressure
	start := State{
		&graph[0],
		make([]*GraphNode[Tunnel], 0),
		0,
	}

	// Build State tree
	DfsCore[*State](0, &start, func(state *State, depth uint) ([]*State, error) {
		children := make([]*State, 0)
		curId := state.Position.Metadata.Id
		// check max depth
		if depth >= MAX_COST {
			if state.PressureReleased > maxPressureReleased {
				maxPressureReleased = state.PressureReleased
				maxState = state
				fmt.Printf("  - Found new max pressure: %d\n", maxPressureReleased)
			}
			return nil, nil
		}
		// check cache
		prev, ok := seen[curId]
		if !ok {
			seen[curId] = make([]uint, MAX_COST)
			seen[curId][depth] = state.PressureReleased
		} else if prev[depth] >= state.PressureReleased {
			return nil, nil
		} else if state.PressureReleased > prev[depth] {
			for i := depth; i < MAX_COST && state.PressureReleased > prev[i]; i++ {
				seen[curId][i] = state.PressureReleased

			}
			// fmt.Printf("Notable Node %c%c @ %d (%d pressure released)\n",
			// curId[0],
			// curId[1],
			// depth, state.PressureReleased,
			// )
		}
		// build move-action children
		for _, ch := range state.Position.To {
			children = append(children, &State{
				ch,
				state.ValvesOpen,
				state.PressureReleased,
			})
		}
		// build open-action child
		isOpen := false
		for _, n := range state.ValvesOpen {
			if n.Metadata.Id == curId {
				isOpen = true
				break
			}
		}
		if !isOpen && state.Position.Metadata.FlowRate > 0 {
			children = append(children, &State{
				state.Position,
				append(state.ValvesOpen, state.Position),
				state.PressureReleased + (MAX_COST-depth-1)*state.Position.Metadata.FlowRate,
			})
		}

		return children, nil
	})

	fmt.Printf("Max Pressure Release Possible: %d\n", maxPressureReleased)
	fmt.Printf("Opened: ")
	for _, n := range maxState.ValvesOpen {
		fmt.Printf("%c%c ", n.Metadata.Id[0], n.Metadata.Id[1])
	}
	fmt.Printf("\n")

	// // Setup Iteration
	// totalPressureRelease := uint(0)
	// tubesOpened := make([]*Tunnel, 0)
	// curNode := &graph[0]
	// curVal := valueNode(curNode)
	// fmt.Printf("Value at start: %d\n", curVal)
	// Iterate
	// for i := 0; i < 4; i++ {
	// printNode(curNode)
	// Find min-cost edge from current node
	// minCost := curVal
	// minNode := curNode.To[0]
	// for _, n := range curNode.To {
	// c := valueNode(n)
	// if c < minCost {
	// minCost = c
	// minNode = n
	// }
	// fmt.Printf("  Cost to %c%c: %d\n", n.Metadata.Id[0], n.Metadata.Id[1], c)
	// }
	// Check if opening current is better
	// curNode.Metadata.Open = true
	// curVal = valueNode(curNode)
	// fmt.Printf("  Cost to open: %d\n", curVal)
	// if curVal < minCost { opening is better
	// tubesOpened = append(tubesOpened, &curNode.Metadata)
	// fmt.Printf("ACTION: Opening %c%c (%d)\n", curNode.Metadata.Id[0], curNode.Metadata.Id[1], curVal)
	// } else { opening is worse, move to min-cost node
	// curNode.Metadata.Open = false
	// curNode = minNode
	// curVal = minCost
	// fmt.Printf("ACTION: Moving to %c%c (%d)\n", curNode.Metadata.Id[0], curNode.Metadata.Id[1], curVal)
	// }
	// Find pressure release
	// for _, tube := range tubesOpened {
	// totalPressureRelease += tube.FlowRate
	// }
	// }
	// Print Results
	// fmt.Printf("Total pressure release: %d\n", totalPressureRelease)
}
