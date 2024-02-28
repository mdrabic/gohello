package slices

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(numbers ...[]int) int {
	sum := 0
	for _, nums := range numbers {
		for _, n := range nums {
			sum += n
		}
	}
	return sum
}
