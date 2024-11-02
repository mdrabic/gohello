package datastructures

import (
	"fmt"
	"testing"
)

func TestPreOrderSearch(t *testing.T) {
	head := BuildBinaryTree()
	expected := []int{1, 2, 4, 5, 3, 6, 7}
	seached := head.preOrderSearch()
	fmt.Println(seached)
	for i := 0; i < len(expected); i++ {
		if expected[i] != seached[i] {
			t.Errorf("expected %d found %d at position %d", expected[i], seached[i], i)
		}
	}
}

func BuildBinaryTree() *BinaryTreeNode[int] {
	head := BinaryTreeNode[int]{value: 1}
	level1_Left := BinaryTreeNode[int]{value: 2}
	level1_Left_level2_Left := BinaryTreeNode[int]{value: 4}
	level1_Left_level2_Right := BinaryTreeNode[int]{value: 5}

	level1_Right := BinaryTreeNode[int]{value: 3}
	level1_Right_Level2_Left := BinaryTreeNode[int]{value: 6}
	level1_Right_Level2_Right := BinaryTreeNode[int]{value: 7}

	//assign nodes
	level1_Left.Left = &level1_Left_level2_Left
	level1_Left.Right = &level1_Left_level2_Right
	level1_Right.Left = &level1_Right_Level2_Left
	level1_Right.Right = &level1_Right_Level2_Right
	head.Left = &level1_Left
	head.Right = &level1_Right

	return &head
}
