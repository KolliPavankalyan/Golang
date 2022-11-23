package main

import "fmt"

func main() {
	var cities []string

	fmt.Println(cities == nil)
	fmt.Printf("%#v\n", cities)

	// initialize a slice

	nums := []int{1, 2, 3, 4}
	fmt.Printf("%#v\n", nums)

	// initialize empty values ny using make attibute

	numbers := make([]int, 2)
	fmt.Printf("%#v\n", numbers)
	numbers[0] = 20
	fmt.Printf("%#v", numbers)

	// initialize a slice using defined type
	type a []string
	b := a{"pavan", "kalyan"}
	fmt.Println(b)

	// initialize a slice using type :2

	type fName []string
	fullName := fName{"pavan", "kalyan", "kolli"}
	fmt.Println(fullName)
}
