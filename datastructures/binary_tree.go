package datastructures

type BinaryTreeNode[T any] struct {
	Left, Right *BinaryTreeNode[T]
	value       T
}

func (head *BinaryTreeNode[T]) preOrderSearch() []T {
	path := make([]T, 0, 10)
	return preOrderWalk(head, path)
}

func preOrderWalk[T any](node *BinaryTreeNode[T], path []T) []T {
	if node == nil {
		return path
	}

	path = append(path, node.value)

	if node.Left != nil {
		path = preOrderWalk(node.Left, path)
	}

	if node.Right != nil {
		path = preOrderWalk(node.Right, path)
	}

	return path
}
