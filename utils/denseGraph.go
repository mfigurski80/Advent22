package utils

import (
	"fmt"
	"math"
)

/**
 * A dense graph is a graph where the edges are stored in a matrix.
 * @param <T> The type of the nodes in the graph
 * @param <D> The type of distance to keep in the graph
 * Extended less-generally by AdjacencyGraph (D=bool) and
 * WeightedGraph (D=uint)
 */
type DenseGraph[T any, D comparable] struct {
	Nodes    []T
	Distance [][]D
}

type AdjacencyGraph[T any] DenseGraph[T, bool]
type WeightedGraph[T any] DenseGraph[T, uint]

func (g *DenseGraph[T, D]) String() string {
	ret := fmt.Sprintf("DenseGraph with %d nodes: ", len(g.Nodes))
	for i, n := range g.Nodes {
		ret += fmt.Sprintf("%d: %v, ", i, n)
	}
	ret += fmt.Sprintf("\n")
	for i, n := range g.Distance {
		ret += fmt.Sprintf("  %-2d: %v\n", i, n)
	}
	return ret[:len(ret)-1]
}

func (g *DenseGraph[T, D]) FilterNodes(f func(T) bool) *DenseGraph[T, D] {
	newNodes := make([]T, 0, len(g.Nodes))
	idMap := make(map[uint]uint)
	for i, node := range g.Nodes {
		keep := f(node)
		if keep {
			newNodes = append(newNodes, node)
			idMap[uint(i)] = uint(len(newNodes) - 1)
		}
	}
	newDistance := make([][]D, len(newNodes))
	for i, row := range g.Distance {
		if _, ok := idMap[uint(i)]; ok {
			newDistance[idMap[uint(i)]] = make([]D, len(newNodes))
			for j, dist := range row {
				if _, ok := idMap[uint(j)]; ok {
					newDistance[idMap[uint(i)]][idMap[uint(j)]] = dist
				}
			}
		}
	}
	g.Nodes = newNodes
	g.Distance = newDistance
	return g
}

func AdjacencyGraphFromSparse[T comparable](g []GraphNode[T]) *DenseGraph[T, bool] {
	nodes := make([]T, len(g))
	nodesIdx := make(map[T]uint)
	adjacency := make([][]bool, len(g))
	for i := range adjacency {
		adjacency[i] = make([]bool, len(g))
	}
	for i, n := range g {
		nodes[i] = n.Metadata
		nodesIdx[n.Metadata] = uint(i)
	}
	for i, n := range g {
		for _, e := range n.To {
			adjacency[i][nodesIdx[e.Metadata]] = true
		}
	}
	return &DenseGraph[T, bool]{nodes, adjacency}
}

func WeightedGraphFromSparse[T comparable](g []GraphNode[T]) *DenseGraph[T, uint] {
	nodes := make([]T, len(g))
	nodesIdx := make(map[T]uint)
	distance := make([][]uint, len(g))
	for i := range distance {
		distance[i] = make([]uint, len(g))
		for j := range distance[i] {
			distance[i][j] = math.MaxUint32
		}
	}
	for i, n := range g {
		nodes[i] = n.Metadata
		nodesIdx[n.Metadata] = uint(i)
	}
	for i, n := range g {
		for _, e := range n.To {
			distance[i][nodesIdx[e.Metadata]] = 1
		}
	}
	// get power of distance matrix
	for k := range distance {
		for i := range distance {
			for j := range distance {
				if distance[i][k]+distance[k][j] < distance[i][j] {
					distance[i][j] = distance[i][k] + distance[k][j]
				}
			}
		}
	}

	return &DenseGraph[T, uint]{nodes, distance}
}
