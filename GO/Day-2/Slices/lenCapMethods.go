package main

import "fmt"

func main() {
	var nums []int
	fmt.Printf("len %d capacity %d\n", len(nums), cap(nums))

	nums = append(nums, 1, 2, 3, 4)
	fmt.Printf("len %d capacity %d\n", len(nums), cap(nums))

	nums = append(nums, 5)
	fmt.Printf("len %d capacity %d\n", len(nums), cap(nums))

	fmt.Println(nums)

}
