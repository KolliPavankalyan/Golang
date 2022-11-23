package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Println("please enter n value")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		var star string
		for j := 1; j <= i; j++ {
			star += strconv.Itoa(j) + " "

		}
		println(star)
	}
}
