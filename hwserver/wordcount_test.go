package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestWordCount(t *testing.T) {
	input := "I am learning Go!"
	output := WordCount(input)
	expected := map[string]int{"Go!": 1, "I": 1, "am": 1, "learning": 1}

	//fmt.Sprint will print the maps keys in key-sorted order
	//could checkout https://github.com/google/go-cmp as a better first class solution
	if fmt.Sprint(output) != fmt.Sprint(expected) {
		t.Errorf("Epected to find needle but got false")
	}
}

func TestFibonacci(t *testing.T) {
	expected := "[0 1 2 3 5 8 13]"
	actual := Fibonacci(7)

	if strings.Compare(expected, actual) != 0 {
		t.Error("Fib strings are not equal")
	}
}
