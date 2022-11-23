package main

import "fmt"

func main() {
	var name string = "pavankalyan"
	fmt.Println(&name)

	//
	var x int = 2
	ptr := &x
	fmt.Printf("the value of ptr %v and type of ptr %T\n", ptr, ptr)

	// address of varible

	p := new(int)

	y := 100

	p = &y

	fmt.Printf("the value of ptr %v type of %T address of variable %p\n", p, p, &p)

	*p = 90
	fmt.Printf("%v", *p)
}
