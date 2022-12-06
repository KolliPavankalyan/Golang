package main

import "fmt"

func linear(Arr []int, n int, Key int) int {
	index := 0
	for index < n {
		if Arr[index] == Key {
			return index
		}
		index += 1
	}
	return -1
}

func main() {
	arr := []int{84, 21, 47, 96, 15}
	n := len(arr)
	key := 100
	fmt.Println(linear(arr, n, key))
}
