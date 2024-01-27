package datastructures 

import "math"

func search(haystack []int, needle int) bool {
	lo := 0
	hi := len(haystack)

	for lo < hi {
		var m_raw float64 = float64(lo + (hi - lo)) / 2.0
		var m = int(math.Floor(m_raw));
		var found = haystack[m];

		if found == needle {
			return true
		} else if found < needle {
			lo = m - 1
		} else {
			hi = m
		}
	}

	return false;
}

func twoCrystalBallsSearch(breaks []int) int {
	//const jumpAmount = math.Floor(math.Sqrt(len(breaks)))

	return -1
}
