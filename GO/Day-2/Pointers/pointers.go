package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(p)
	x := 100
	p = &x
	fmt.Printf("the value of p is  %v and value of x is  %v\n", p, x)

	*p = 200
	fmt.Printf("the value of x is %v and value of p is %v\n", x, p)
	str := fmt.Sprint("%v", 100)
	fmt.Println(str)
}
