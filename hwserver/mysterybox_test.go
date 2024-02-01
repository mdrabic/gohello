package main

import (
	"log"
	"testing"
)

func TestSurpriseReducesBoxCount(t *testing.T) {
	mysteryBoxSize := count()
	if mysteryBoxSize < 1 {
		for i := 0; i < 2; i++ {
			put(42)
			put(68)
			put(96)
		}
	}

	preCountSurprise := count()
	surprise()
	postCountSurprise := count()

	if postCountSurprise >= preCountSurprise {
		log.Println("found: ", preCountSurprise, postCountSurprise, 3)
		t.Errorf("Epected to find needle but got false")
	}
}
