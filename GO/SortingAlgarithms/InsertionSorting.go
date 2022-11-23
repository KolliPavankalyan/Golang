package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	piece := createpiece(20)
	fmt.Println("\n Unsorted \n", piece)
	Insertionsort(piece)
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

func Insertionsort(items []int) {
	n := len(items)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
		}
	}
}
