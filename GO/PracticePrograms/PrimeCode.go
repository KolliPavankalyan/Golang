package main

import "fmt"

func main() {
	var value int
	value = 100
	for i := 2; i <= value; i++ {
		var prime int

		for j := 2; j < i; j++ {
			if (i % j) == 0 {
				break
			} else {
				prime = i
			}
		}
		if prime != 0 {
			fmt.Println(prime)
		}
	}
}
