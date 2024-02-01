package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	tokens := strings.Fields(s)
	m := make(map[string]int)

	for _, value := range tokens {
		e, ok := m[value]
		if ok {
			m[value] = e + 1
		} else {
			m[value] = 1
		}
	}
	fmt.Println("m: ", m)
	return m
}
