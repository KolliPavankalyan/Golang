package main

import (
	"fmt"
	"strconv"
)

type linkelist struct {
	number int
	next   *linkelist
}

// insert a linkedlist at end
func (node *linkelist) insert(num int) {
	var temp = &linkelist{}
	temp.number = num
	temp.next = nil
	if node == nil {
		node = temp
	} else {
		var p = &linkelist{}
		p = node
		for p.next != nil {
			p = p.next
		}
		p.next = temp
	}
}

// display linked list
func (node *linkelist) display() {
	var p = &linkelist{}
	p = node.next
	for p != nil {
		fmt.Printf("%d ->", p.number)
		p = p.next
	}

}
func (node *linkelist) DeleteFromDirst() *linkelist {
	node = node.next
	return node
}

// delete from last in linkedlist
func (node *linkelist) DeleteFromLast() {
	var p *linkelist = node
	for p.next.next != nil {
		p = p.next
	}
	p.next = nil
}

func main() {
	head := new(linkelist)
	var choice string

	for true {
		fmt.Println("\n  enter choice")
		fmt.Println("1. Insert the linkedlist")
		fmt.Println("2. display linkedlist")
		fmt.Println("3. Delete from begining")
		fmt.Println("4. Deleting from last")
		fmt.Scanln(&choice)
		var data string
		switch choice {
		case "1":
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			head.insert(num)
		case "2":
			head.display()
		case "3":
			head.DeleteFromDirst()
		case "4":
			head.DeleteFromLast()
		}
	}
}
