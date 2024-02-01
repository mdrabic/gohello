package main

import (
	"fmt"
	"testing"
)

func TestWordCount(t *testing.T) {
	input := "I am learning Go!"
	output := WordCount(input)
	expected := map[string]int{"Go!": 1, "I": 1, "am": 1, "learning": 1}

	//fmt.Sprint will print the maps keys in key-sorted order
	if fmt.Sprint(output) != fmt.Sprint(expected) {
		t.Errorf("Epected to find needle but got false")
	}
}
