package main

import "fmt"

func reverse(a []int) []int {
	length := len(a)
	arr := []int{}
	for i := length - 1; i >= 0; i-- {
		arr = append(arr, a[i])

	}
	return arr
}

func main() {
	a := []int{1, 2, 3}

	fmt.Println(reverse(a))
}
