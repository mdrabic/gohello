package datastructures

import (
	"fmt"
	"testing"
)

func TestBinary(t *testing.T) {
	nums := []int{2, 4, 42, 69, 420, 690, 4242}
	needle := 4
	result := binary(nums, needle)
	if !result {
		t.Errorf("Epected to find needle but got false")
	}
}

func TestBubble(t *testing.T) {
	nums := []int{4242, 690, 69, 42, 4, 2}
	bubble(nums)
	fmt.Println(nums)
	if nums[0] != 2 {
		t.Errorf("Epected to find 2 at position 0")
	}
}

func TestBubbleEmptySlice(t *testing.T) {
	nums := []int{}
	bubble(nums)
	fmt.Println(nums)
	if len(nums) != 0 {
		t.Errorf("Expected slice to be empty")
	}
}

func TestBubbleOddSize(t *testing.T) {
	nums := []int{4242, 690, 69, 42, 2}
	bubble(nums)
	fmt.Println(nums)
	if nums[0] != 2 {
		t.Errorf("Epected to find 2 at position 0")
	}
}
