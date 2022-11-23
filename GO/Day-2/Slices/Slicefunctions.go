package main

import (
	"fmt"
	"strings"
)

func main() {

	//slice function

	s1 := []int{1, 2, 3, 4, 5}
	s3 := s1[2:5]
	s2 := s1[0:3]
	//s3[0] = 10

	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Println(len(s3), cap(s3))
	fmt.Println(len(s2), cap(s2))

	//arayys slice
	fmt.Println(strings.Repeat("#", 10))
	a1 := [5]int{1, 2, 3, 4, 5}
	a2 := a1[2:5]
	a3 := a1[0:3]
	a2[2] = 20

	fmt.Println(a2)
	fmt.Println(a3)

	fmt.Println(len(a2), cap(a2))
	fmt.Println(len(a3), cap(a3))

}
