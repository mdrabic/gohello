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

func Fibonacci(numbers int) string {
	f := fib()
	sequence := make([]int, 0, numbers)

	for i := 0; i < numbers; i++ {
		n := f()
		fmt.Println(n)
		sequence = append(sequence, n)
	}

	return fmt.Sprint(sequence)
}

func fib() func() int {
	a, b := 0, 0
	return func() int {
		c := a + b
		a = b

		if c <= 0 {
			b = 1
		} else {
			b = c
		}

		return c
	}
}
