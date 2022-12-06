package main

import "fmt"

func BinarySearch(Arr []int, n, k int) int {
	left := 0
	right := n - 1

	for left <= right {
		median := (left + right) / 2
		if Arr[median] == k {
			return median
		} else if k < Arr[median] {
			right = median - 1
		} else if median < k {
			left = median + 1
		}
	}
	return -1
}

func main() {
	slice := []int{1, 20, 50, 70}
	n := len(slice)
	k := 50
	fmt.Println(BinarySearch(slice, n, k))
}
