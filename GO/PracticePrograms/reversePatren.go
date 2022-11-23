package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Println("enter n value")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		var star string
		for j := 1; j <= i; j++ {
			star += strconv.Itoa(i+1-j) + " "
		}
		fmt.Println(star)
	}
}
