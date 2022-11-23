package main

import "fmt"

func main() {
	nums := []int{1, 2, 1}
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}
	fmt.Println(nums[0])
}
