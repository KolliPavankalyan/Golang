package main

import "fmt"

func main() {
	var n int
	fmt.Println("please enter n value")
	fmt.Scanln(&n)
	var arr = make([]int, n)
	for i := 0; i <= n; i++ {
		fmt.Println("enter array values")
		fmt.Scanf("%d", &arr[i])

	}
	fmt.Println(arr)
}
