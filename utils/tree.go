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
