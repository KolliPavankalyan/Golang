package main

import "fmt"

func f1(a ...int) {
	fmt.Println(a)

}
func main() {
	f1(1, 2, 3, 4, 56)
	nums := []int{1, 2}
	nums = append(nums, 4, 5, 6)
	f1(nums...)
	fmt.Println(len(nums))
}
