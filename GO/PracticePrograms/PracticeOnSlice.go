package main

import "fmt"

func main() {
	var name = []int{1, 2, 3, 4, 5}
	sum := 0
	for i, v := range name {
		_ = v
		sum += name[i]
	}
	fmt.Println(sum)

}
