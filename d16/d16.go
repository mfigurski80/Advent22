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

type State struct {
	Position         uint
	Depth            uint
	PressureReleased uint
	ValvesOpen       []uint
}

func (s *State) String() string {
	return fmt.Sprintf("<State: (%d) depth=%-2d, pressure=%d, valves=%v>", s.Position, s.Depth, s.PressureReleased, s.ValvesOpen)
}

func Main() {
	// Build Graph
	graph := buildTunnelGraphFromFile("d16/in.txt")
	dg := WeightedGraphFromSparse(graph).FilterNodes(func(n Tunnel) bool {
		return n.FlowRate > 0 || (n.Id[0] == 'A' && n.Id[1] == 'A')
	})
	fmt.Println(dg)

	// Setup State Representation
	seen := map[TunnelId][]uint{} // cache: tunnel + depth => pressure
	var maxState *State
	maxPressureReleased := uint(0)
	start := State{0, 0, 0, []uint{}}
	fmt.Printf("Starting at %c%c\n", dg.Nodes[start.Position].Id[0], dg.Nodes[start.Position].Id[1])

	// Traverse State tree
	DfsCore[*State](0, &start, func(state *State, _ uint) ([]*State, error) {
		// fmt.Println(state.String())
		// compare to maximum
		if state.PressureReleased > maxPressureReleased {
			maxPressureReleased = state.PressureReleased
			maxState = state
			fmt.Printf("  - Found new max pressure: (%d) ", maxPressureReleased)
			for _, v := range maxState.ValvesOpen {
				fmt.Printf("%c%c ", dg.Nodes[v].Id[0], dg.Nodes[v].Id[1])
			}
			fmt.Printf("\n")
		}
		// check max depth
		if state.Depth >= MAX_COST || len(state.ValvesOpen) >= len(dg.Nodes) {
			return nil, nil
		}
		// check cache
		curId := dg.Nodes[state.Position].Id
		_, ok := seen[curId]
		if !ok {
			seen[curId] = make([]uint, MAX_COST)
		}
		if ok && seen[curId][state.Depth] >= state.PressureReleased {
			return nil, nil // cached is better
		} else if state.PressureReleased > seen[curId][state.Depth] {
			for i := state.Depth; i < MAX_COST && state.PressureReleased > seen[curId][i]; i++ {
				seen[curId][i] = state.PressureReleased
			}
			// fmt.Printf("Notable Node %c%c @ %d (%d pressure released)\n", curId[0], curId[1], state.Depth, state.PressureReleased)
		}
		children := make([]*State, 0)
	CONTINUE:
		for i, d := range dg.Distance[state.Position] {
			i := uint(i)
			if i == state.Position {
				continue // skip self
			}
			for _, n := range state.ValvesOpen {
				if i == n {
					continue CONTINUE // skip if already open
				}
			}
			expDepth := state.Depth + MOVE_COST*d + OPEN_COST
			if expDepth > MAX_COST {
				continue // skip if too far
			}
			child := State{
				i,
				expDepth,
				state.PressureReleased + (MAX_COST-expDepth)*dg.Nodes[i].FlowRate,
				make([]uint, len(state.ValvesOpen)+1),
			}
			// DO NOT USE APPEND: it reuses the same underlying array
			copy(child.ValvesOpen, state.ValvesOpen)
			child.ValvesOpen[len(child.ValvesOpen)-1] = i
			children = append(children, &child)
		}
		return children, nil
	})

	fmt.Printf("Max Pressure Release Possible: %d\n", maxPressureReleased)
	if maxState != nil {
		fmt.Printf("Opened: ")
		for _, n := range maxState.ValvesOpen {
			fmt.Printf("%c%c ", dg.Nodes[n].Id[0], dg.Nodes[n].Id[1])
		}
		fmt.Printf("\n")
	}
}
