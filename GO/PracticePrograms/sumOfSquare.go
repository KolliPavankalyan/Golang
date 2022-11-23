package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	n := 5
	var v float64
	var total float64
	for i := 1; i < n; i++ {
		v = float64(i)
		total += (math.Pow(v, 2))
	}
	fmt.Println(total)
	// normal integer
	fmt.Println(strings.Repeat("-", 20))
	res := math.Pow(3, 3)
	fmt.Println(res)

}
