package main

import (
	"fmt"
	"math"
)

func f1() {
	fmt.Println("this function")
}
func f2(a int, b int) {
	fmt.Println(a + b)
}

func f3(a float64) {
	fmt.Println(math.Pow(a, a))
}

func main() {
	f1()
	f2(10, 20)
	f3(1.2)
}
