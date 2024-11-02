package datastructures

import (
	"testing"
)

func TestPush(t *testing.T) {
	linkedList := LinkedList[string]{Length: 0}
	linkedList.push("hello")

	if linkedList.Length != 1 {
		t.Errorf("Epected length of 1")
	}
}

func TestPeek(t *testing.T) {
	linkedList := LinkedList[string]{Length: 0}
	linkedList.push("hello")

	value := linkedList.peek()
	if value != "hello" {
		t.Errorf("Expected 'hello' but got 'value'")
	}
}

func TestPopOneItem(t *testing.T) {
	linkedList := LinkedList[string]{Length: 0}
	linkedList.push("hello")
	linkedList.push("world")

	value := linkedList.pop()
	if value != "hello" {
		t.Errorf("Expected 'hello' but got 'value'")
	}

	head := linkedList.peek()
	if head != "world" {
		t.Errorf("Expected 'world' but got 'value'")
	}
}
