package main

import "fmt"

type sli []int

// is empty
func (s *sli) isEmpty() bool {
	if len(*s) == 0 {
		return true
	} else {
		return false
	}
}

//adding to stack

func (s *sli) append(val int) {
	*s = append(*s, val)
}

// remove element in stack
func (s *sli) pop() {
	if len(*s) == 0 {
		fmt.Println("No element is found in stack")
	}
	index := len(*s) - 1
	*s = (*s)[:index]

}

//display stack
func (s *sli) display() {
	for _, v := range *s {
		fmt.Println(v)
	}
	fmt.Println()
}

func main() {
	var slice sli
	slice.append(10)
	slice.display()
	slice.append(20)
	slice.append(30)
	slice.display()
	slice.pop()
	slice.display()

}
