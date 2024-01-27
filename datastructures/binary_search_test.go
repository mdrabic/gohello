package datastructures 

import "testing"

func TestSearch(t *testing.T)  {
	nums := []int{2,4,42,69,420,690,4242}
	needle := 4
	result := search(nums, needle)
	if (!result) {
		t.Errorf("Epected to find needle but got false")
	}
}

