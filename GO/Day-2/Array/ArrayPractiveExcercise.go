package main

import "fmt"

func main() {
	//decleare array
	var cities [2]string
	fmt.Println(cities)

	//decleare array and intialize array values
	grades := [3]float64{20.2, 5.5}
	fmt.Println(grades)
	//ellipsis operator
	taskDone := [...]bool{true, false, true}
	fmt.Println(taskDone)
}
