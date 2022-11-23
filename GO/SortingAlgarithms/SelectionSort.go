package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	piece := createpiece(20)
	fmt.Println("\n Unsorted \n", piece)
	selectionsort(piece)
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

func selectionsort(items []int) {
	length := len(items)
	for i := 0; i < length; i++ {
		var min = i
		for j := i; j < length; j++ {
			if items[j] < items[min] {
				min = j
			}
		}
		items[i], items[min] = items[min], items[i]
	}
}
