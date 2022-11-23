package main

import (
	"fmt"
	"strconv"
)

func main() {
	//s := string(100)
	//fmt.Println(s)

	// convert num to string
	s1 := "200"
	_ = s1
	var myStr = fmt.Sprintf("%d", 425)
	fmt.Println(myStr)

	// convert string to num
	s1 = "3.01"

	var f1, err = strconv.ParseFloat(s1, 64)
	_ = err
	fmt.Println(f1)
	fmt.Printf("%T", f1)
}
