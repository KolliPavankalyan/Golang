package main

import "fmt"

func main() {
	var x int = 3
	var y = 3.1

	x = int(float64(y))
	fmt.Println(x)
}
