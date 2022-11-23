package main

import "fmt"

func main() {
	var p1 *int

	p2 := 10

	p1 = &p2
	fmt.Printf("value of a p1 is %v", p1)
	fmt.Printf("Address of p1 %v", &p1)

}
