package slices

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	a := []int{10, 20, 30}
	b := []int{1, 2, 3}
	got := SumAll(a, b)
	want := 66
	if got != want {
		t.Errorf("got %d want %d given, %v %v", got, want, a, b)
	}
}
