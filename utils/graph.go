package utils

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
	return BfsCore(root, func(node *GraphNode[T], level uint) ([]*GraphNode[T], error) {
		err := f(node, level)
		if err != nil {
			return nil, err
		}
		children := []*GraphNode[T]{}
		for i := range node.To {
			if node.To[i].VisitedDepth > 0 {
				continue
			}
			node.To[i].VisitedDepth = level + 1
			children = append(children, node.To[i])
		}
		return children, nil
	})
}
