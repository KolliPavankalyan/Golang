package main

func main() {
	//const a = "hello" + "world"
	//fmt.Println("%T\n", a)

	const (
		a = iota
		b = iota
		c = iota
	)
	println(a, b, c)
}
