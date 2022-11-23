package main

import (
	"fmt"
	"math"
)

func main() {
	var a int8 = 127
	fmt.Printf("%T", a)
	f := float32(math.MaxFloat32)
	fmt.Println(f)
}
