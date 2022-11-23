package main

import "fmt"

func main() {
	b := 100
	for a := 0; a <= b; a++ {
		if a%2 == 0 {
			fmt.Printf("%d\n", a)
		}
	}
}
