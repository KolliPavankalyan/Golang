package main

import "fmt"

func main() {
	type book struct {
		title string
		year  int
	}
	type book1 struct {
		title string
		year  int
	}

	mybook := book{title: "pavan"}
	fmt.Printf("%#v", mybook)

	mybook1 := book1{title: "rebel"}
	fmt.Printf("%#v", mybook1)
}
