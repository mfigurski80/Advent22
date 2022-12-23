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
	Open     bool
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
					Open:     false,
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

func valueNode(n *GraphNode[Tunnel]) uint {
	val := uint(0)
	n.BfsOnGraph(0, func(n *GraphNode[Tunnel], depth uint) error {
		// printNode(n)
		if n.Metadata.Open { // no cost if already open
			return nil
		}
		val += n.Metadata.FlowRate * (MOVE_COST*depth + OPEN_COST)
		return nil
	})
	return val
}

func Main() {
	// Build Graph
	graph := buildTunnelGraphFromFile("d16/in.txt")
	if len(graph) > 256 {
		panic("Misjudged graph size")
	}
	// Ensure Okay
	for i := range graph {
		printNode(&graph[i])
	}

	// Setup Iteration
	totalPressureRelease := uint(0)
	tubesOpened := make([]*Tunnel, 0)
	curNode := &graph[0]
	curVal := valueNode(curNode)
	fmt.Printf("Value at start: %d\n", curVal)
	// Iterate
	for i := 0; i < 4; i++ {
		printNode(curNode)
		// Find min-cost edge from current node
		minCost := curVal
		minNode := curNode.To[0]
		for _, n := range curNode.To {
			c := valueNode(n)
			if c < minCost {
				minCost = c
				minNode = n
			}
			fmt.Printf("  Cost to %c%c: %d\n", n.Metadata.Id[0], n.Metadata.Id[1], c)
		}
		// Check if opening current is better
		curNode.Metadata.Open = true
		curVal = valueNode(curNode)
		fmt.Printf("  Cost to open: %d\n", curVal)
		if curVal < minCost { // opening is better
			tubesOpened = append(tubesOpened, &curNode.Metadata)
			fmt.Printf("ACTION: Opening %c%c (%d)\n", curNode.Metadata.Id[0], curNode.Metadata.Id[1], curVal)
		} else { // opening is worse, move to min-cost node
			curNode.Metadata.Open = false
			curNode = minNode
			curVal = minCost
			fmt.Printf("ACTION: Moving to %c%c (%d)\n", curNode.Metadata.Id[0], curNode.Metadata.Id[1], curVal)
		}
		// Find pressure release
		for _, tube := range tubesOpened {
			totalPressureRelease += tube.FlowRate
		}
	}
	// Print Results
	fmt.Printf("Total pressure release: %d\n", totalPressureRelease)
}
