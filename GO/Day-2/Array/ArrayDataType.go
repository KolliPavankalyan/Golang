package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers [4]int

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	array_name := [4]string{"item1", "item2", "item3", "itemN"}
	fmt.Printf("%#v\n", array_name)

	floatVariable := [...]float32{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}
	fmt.Printf("%#v\n", floatVariable)

	intArray := [...]int{1,
		2,
		3,
		4,
		5,
	}
	fmt.Printf("%#v", intArray)

	//modify the array
	var Marray = [5]int{1, 2, 3, 4, 5}

	for i, v := range Marray {
		fmt.Println("index:", i, "value:", v)
	}

	for i := 0; i <= len(numbers); i++ {
		fmt.Println("index:", i, "value:", Marray[i])
	}

	// repeat number of times one item

	fmt.Printf(strings.Repeat("8", 10))

	// inner arrays

	students := [3][2]int{
		{1, 2},
		{4, 5},
		{7, 8},
	}

	fmt.Println(students)

	//check arrays

	a1 := [4]int{1, 2, 3, 4}
	a2 := a1

	fmt.Println("equel r not", a1 == a2)
	a1[0] = 5
	fmt.Println("equel r not", a1 == a2)

	a2 = a1
	fmt.Println("equel r not", a1 == a2)

	//arrays with keyed elements

	cities := [...]string{
		1: "karimnagar",
		"Nellore",
		5: "ts"}
	fmt.Println(cities)

}
