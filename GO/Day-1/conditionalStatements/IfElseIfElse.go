package main

import "fmt"

func main() {
	// find given number is even or odd or zero
	num := 0
	if num > 0 {
		fmt.Println("even")
	} else if num < 0 {
		fmt.Println("odd")
	} else {
		fmt.Println("zero")
	}

}
