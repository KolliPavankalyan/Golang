package main

import "fmt"

func main() {
	isPalindrome(121)
}

func isPalindrome(x int) {
	input_num := x
	var remainder int
	res := 0
	for x > 0 {
		remainder = x % 10
		res = (res * 10) + remainder
		fmt.Println(res)
		x = x / 10
	}
	if input_num == res {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
