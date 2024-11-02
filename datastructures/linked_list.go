package datastructures

import ()

type Node[T any] struct {
	Prev, Next *Node[T]
	value      T
}

type LinkedList[T any] struct {
	Head   *Node[T]
	Tail   *Node[T]
	Length int
}

func (ll *LinkedList[T]) push(n T) {
	node := Node[T]{value: n}

	if ll.Head == nil {
		ll.Head = &node
		ll.Tail = &node
	} else {
		ll.Tail.Next = &node
		ll.Tail = &node
	}

	ll.Length += 1
}

func (ll *LinkedList[T]) length() int {
	return ll.Length
}

func (ll *LinkedList[T]) isEmpty() bool {
	return ll.length() == 0
}

func (ll *LinkedList[T]) peek() T {
	return ll.Head.value
}

func (ll *LinkedList[T]) pop() T {
	value := ll.Head.value

	if !ll.isEmpty() {
		ll.Length -= 1
		ll.Head = ll.Head.Next
	} else {
		ll.Head = nil
		ll.Tail = nil
	}

	return value
}
