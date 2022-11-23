package main

import (
	"fmt"
)

func change(x *int) *float64 {
	*x = 100
	b := 5.5
	return &b
}
func main() {
	a := 9
	p := &a
	fmt.Printf("value a is before changes %v\n", a)
	change(p)
	fmt.Printf("value a is before changes %v", a)

}
