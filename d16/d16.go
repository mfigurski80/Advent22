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

type SearchState = []uint

func PathCost(g *DenseGraph[Tunnel, uint], s SearchState, start uint) uint {
	pressure := uint(0)
	dist := uint(0)
	curId := start
	for _, nextId := range s {
		addedDist := g.Distance[curId][nextId]*MOVE_COST + OPEN_COST
		if dist+addedDist >= MAX_COST {
			break
		}
		dist += addedDist
		pressure += g.Nodes[nextId].FlowRate * (MAX_COST - dist)
	}
	return pressure
}

func Main() {
	// Build Graph
	graph := buildTunnelGraphFromFile("d16/in.txt")
	dg := WeightedGraphFromSparse(graph).FilterNodes(func(n Tunnel) bool {
		return n.FlowRate > 0 || (n.Id[0] == 'A' && n.Id[1] == 'A')
	})
	fmt.Println(dg)

	// Setup State
	maxPressure := uint(0)
	maxState := SearchState{}
	DoForPermutations(uint(len(dg.Nodes)), func(n uint, s SearchState) {
		fmt.Println(s)
		pressure := PathCost(dg, s, 0)
		if pressure > maxPressure {
			maxPressure = pressure
			maxState = s
		}
	})

	// Show Results
	fmt.Printf("Max Pressure Release Possible: %d\n", maxPressure)
	if maxState != nil {
		fmt.Printf("Opened: ")
		for _, n := range maxState {
			fmt.Printf("%c%c ", dg.Nodes[n].Id[0], dg.Nodes[n].Id[1])
		}
		fmt.Printf("\n")
	}
}
