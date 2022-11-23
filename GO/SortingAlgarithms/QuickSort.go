package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	piece := createpiece(10)
	fmt.Println("\n Unsorted \n", piece)
	Quicksort(piece)
	fmt.Println("\n sorted \n", piece)
}

func createpiece(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)

	}
	return slice
}

func Quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	center := rand.Int() % len(a)
	a[center], a[right] = a[right], a[center]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]

	Quicksort((a[:left]))
	Quicksort(a[left+1:])
	return a
}
