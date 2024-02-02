package main

import "math/rand"

var box = make([]int, 0, 10)

func put(item int) []int {
	box = append(box, item)
	return box
}

func count() int {
	return len(box)
}

func surprise() int {
	i := rand.Intn(len(box))
	item := box[i]
	box = append(box[:i], box[i+1:]...)
	return item
}
