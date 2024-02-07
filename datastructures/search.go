package datastructures

import "math"

func binary(haystack []int, needle int) bool {
	lo := 0
	hi := len(haystack)

	for lo < hi {
		var m_raw float64 = float64(lo+(hi-lo)) / 2.0
		var m = int(math.Floor(m_raw))
		var found = haystack[m]

		if found == needle {
			return true
		} else if found < needle {
			lo = m - 1
		} else {
			hi = m
		}
	}

	return false
}

func bubble(arr []int) {
	var swap int
	for i := 0; i < len(arr)-1; i++ {
		for a, b := 0, 1; b < len(arr)-i; a, b = a+1, b+1 {
			if arr[a] > arr[b] {
				swap = arr[b]
				arr[b] = arr[a]
				arr[a] = swap
			}
		}
	}
}

func twoCrystalBallsSearch(breaks []int) int {
	//const jumpAmount = math.Floor(math.Sqrt(len(breaks)))

	return -1
}
