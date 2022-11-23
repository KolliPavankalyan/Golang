package main

import "fmt"

func main() {
	var employee map[string]string

	fmt.Printf("%#v\n", (employee))

	// decleare a map

	var members map[string]string

	fmt.Printf("%q\n", members["name"])

	// map1
	map1 := make((map[int]string))
	map1[0] = "pavan"
	fmt.Println(map1)

	v, ok := map1[0]
	if ok {
		fmt.Println("rom balance is", v)

	} else {
		fmt.Println("ok is false")
	}

}
