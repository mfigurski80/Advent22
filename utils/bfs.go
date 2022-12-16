package utils

type BfsDecisionFunction[T any] func(T, uint) ([]T, error)

func BfsCore[T any](root T, f BfsDecisionFunction[T]) error {
	queue := Queue[T]()
	queue.Push(root)
	level := uint(0)
	nextLevelAt := uint(1)
	for queue.Length() > 0 {
		node := queue.Pop()
		children, err := f(node, level)
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
