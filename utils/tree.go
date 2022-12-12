package utils

type TreeNode[T any] struct {
	Children []*TreeNode[T]
	Parent   *TreeNode[T]
	Metadata T
}

func (root TreeNode[T]) DfsOnTree(level int, f func(TreeNode[T], int)) {
	f(root, level)
	for _, child := range root.Children {
		child.DfsOnTree(level+1, f)
	}
}
