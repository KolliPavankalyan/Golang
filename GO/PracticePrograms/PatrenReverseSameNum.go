package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Println("enter n value")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		num := n + 1 - i
		println(strings.Repeat(strconv.Itoa(num), i), "  ")
	}
}
