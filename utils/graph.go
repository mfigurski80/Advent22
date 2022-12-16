package utils

type GraphNode[T any] struct {
	To           []*GraphNode[T]
	From         []*GraphNode[T]
	Metadata     T
	VisitedDepth uint32
}

func (root *GraphNode[T]) DfsOnGraph(level uint, f func(*GraphNode[T], uint)) {
	f(root, level)
	for i := range root.To {
		child := root.To[i]
		if child.VisitedDepth > 0 && child.VisitedDepth < uint32(level) {
			continue
		}
		child.VisitedDepth = uint32(level)
		child.DfsOnGraph(level+1, f)
	}
}

func (root *GraphNode[T]) BfsOnGraph(level uint, f func(*GraphNode[T], uint) error) error {
	queue := Queue[*GraphNode[T]]()
	root.VisitedDepth = uint32(level)
	queue.Push(root)
	for queue.Length() > 0 {
		node := queue.Pop()
		err := f(node, uint(node.VisitedDepth))
		if err != nil {
			return err
		}
		for i := range node.To {
			if node.To[i].VisitedDepth > 0 {
				continue
			}
			node.To[i].VisitedDepth = node.VisitedDepth + 1
			queue.Push(node.To[i])
		}
	}
	return nil
}
