package main

import "fmt"

type sli []int

// is Empty r not
func (s *sli) isEmpty() bool {
	if len(*s) == 0 {
		return true
	} else {
		return false
	}
}

// add to queue
func (s *sli) Enqueue(val int) {
	*s = append(*s, val)
}

// display queue
func (s *sli) display() {
	for _, v := range *s {
		fmt.Println(v)
	}
	println()
}

//delete in queue
func (s *sli) Dequeue() {
	*s = (*s)[1:]
}

func main() {
	var slice sli
	slice.Enqueue(10)
	slice.Enqueue(20)
	slice.display()
	slice.Dequeue()
	slice.display()
}
