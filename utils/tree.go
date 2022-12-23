package utils

type TreeNode[T any] struct {
	Children []*TreeNode[T]
	Parent   *TreeNode[T]
	Metadata T
}

func (root TreeNode[T]) DfsOnTree(depth int, f func(TreeNode[T], int)) {
	f(root, depth)
	for _, child := range root.Children {
		child.DfsOnTree(depth+1, f)
	}
}

func (root *TreeNode[T]) BfsOnTree(depth int, f func(*TreeNode[T], int)) error {
	return BfsCore(root, func(node *TreeNode[T], depth uint) ([]*TreeNode[T], error) {
		f(node, int(depth))
		return node.Children, nil
	})
}

func ConstructTree[T any](depth uint, root T, f func(T, uint) []T) *TreeNode[T] {
	tree := &TreeNode[T]{Metadata: root}
	children := f(root, depth)
	for _, child := range children {
		tree.Children = append(tree.Children, ConstructTree(depth+1, child, f))
	}
	return tree
}
