package utils

import "sort"

/**
 * Decision function for BFS
 * @param node: current node
 * @param level: current level
 * @return: children of the current node
 */
type BfsDecisionFunction[T any] func(T, uint) ([]T, error)

func BfsCore[T any](root T, fn BfsDecisionFunction[T]) error {
	queue := Queue[T]()
	queue.Push(root)
	level := uint(0)
	nextLevelAt := uint(1)
	for queue.Length() > 0 {
		node := queue.Pop()
		children, err := fn(node, level)
		if err != nil {
			return err
		}
		for i := range children {
			queue.Push(children[i])
		}
		nextLevelAt--
		if nextLevelAt == 0 {
			level++
			nextLevelAt = uint(queue.Length())
		}
	}
	return nil
}

/**
 * DFS Search Decision Function
 * @param node: current node
 * @param level: current level
 * @return: children of current node
 */
type DfsDecisionFunction[T any] func(T, uint) ([]T, error)

func DfsCore[T any](level uint, root T, fn DfsDecisionFunction[T]) error {
	children, err := fn(root, level)
	if err != nil {
		return err
	}
	for i := range children {
		if err := DfsCore[T](level+1, children[i], fn); err != nil {
			return err
		}
	}
	return nil
}

/**
 * Beam Search Decision Function
 * @param node: current node
 * @param level: current level
 * @return: children of the current node
 * @return: values of children of the current node
 */
type BeamSearchDecisionFunc[T any] func(T, uint) ([]T, []int, error)

func BeamSearchCore[T any](root T, nToKeep uint, fn BeamSearchDecisionFunc[T]) error {
	queue := Queue[T]()
	queue.Push(root)
	level := uint(0)
	nextLevelAt := uint(1)
	for queue.Length() > 0 {
		node := queue.Pop()
		children, values, err := fn(node, level)
		if err != nil {
			return err
		}
		sort.Slice(children, func(i, j int) bool {
			return values[i] < values[j]
		})
		for i := range children[:nToKeep] {
			queue.Push(children[i])
		}
		nextLevelAt--
		if nextLevelAt == 0 {
			level++
			nextLevelAt = uint(queue.Length())
		}
	}
	return nil
}
