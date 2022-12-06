package main

import "fmt"

func BinaryRecursive(Arr []int, key, L, R int) int {
	if L > R {
		return -1
	} else {
		mid := (L + R) / 2
		if key == Arr[mid] {
			return mid
		} else if key < Arr[key] {
			return BinaryRecursive(Arr, key, L, mid-1)
		} else if key > Arr[key] {
			return BinaryRecursive(Arr, key, mid+1, R)
		}

	}
	return -1
}

func main() {
	Arr := []int{10, 25, 35, 47, 95}
	key := 35
	L := 0
	R := len(Arr) - 1
	fmt.Println(Arr, key, L, R)
}
