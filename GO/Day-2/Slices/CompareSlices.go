package main

import "fmt"

func main() {

	var a []string
	fmt.Println(a == nil)  // true
	fmt.Printf("%#v\n", a) //[]string{nil}
	/*fmt.Println(len(a))
	fmt.Println(cap(a))*/

	var b = []string{}
	fmt.Println(b == nil)  //false
	fmt.Printf("%#v\n", b) //[]string{}
	/*fmt.Println(len(b))
	fmt.Println(cap(b))*/

	new_slice := []int{1, 2, 3, 4, 5}
	sliceB := new_slice[1:5]

	sliceA := new_slice[0:3]

	fmt.Println(len(sliceA), cap(sliceA))

	fmt.Println(len(sliceB), cap(sliceB))

}
