package utils

/**
 * Graph Node data structure
 * @param T: type of the data stored in the node
 * @param I: type of the *unique* id of the node
 */
type GraphNode[T any] struct {
	To           []*GraphNode[T]
	From         []*GraphNode[T]
	Metadata     T
	VisitedDepth uint
}

func (root *GraphNode[T]) DfsOnGraph(level uint, f func(*GraphNode[T], uint)) {
	f(root, level)
	for i := range root.To {
		child := root.To[i]
		if child.VisitedDepth > 0 && child.VisitedDepth < level {
			continue
		}
		child.VisitedDepth = level
		child.DfsOnGraph(level+1, f)
	}
}

func (root *GraphNode[T]) BfsOnGraph(level uint, f func(*GraphNode[T], uint) error) error {
	seen := map[*GraphNode[T]]bool{root: true}
	return BfsCore(root, func(node *GraphNode[T], level uint) ([]*GraphNode[T], error) {
		err := f(node, level)
		if err != nil {
			return nil, err
		}
		children := []*GraphNode[T]{}
		for i, ch := range node.To {
			if seen[ch] {
				continue
			}
			seen[ch] = true
			node.To[i].VisitedDepth = level + 1     // mutate
			children = append(children, node.To[i]) // keep reference
		}
		return children, nil
	})
}
