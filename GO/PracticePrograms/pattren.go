package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 5

	for i := 0; i <= n; i++ {
		var star string
		for j := 0; j <= i; j++ {
			p := strconv.Itoa(j)
			star += p + " "
		}
		fmt.Println(star)
	}
}
