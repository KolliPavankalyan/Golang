package main

import (
	"fmt"
	"os"
	"strconv"
)

// creating a struct
type linkelist struct {
	number int
	next   *linkelist //next node store here
}

// Insert Method
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

// display method
func (node *linkelist) display() {
	var p = &linkelist{}
	p = node.next
	for p != nil {
		fmt.Printf("%d ->", p.number)
		p = p.next
	}
}

// delete from first
func (node *linkelist) deletefirst() *linkelist {
	node = node.next
	return node
}

// Delete from last
func (node *linkelist) deletelast() {
	var p *linkelist = node
	for p.next.next != nil {
		p = p.next
	}
	p.next = nil
}

// insert specific position
func (node *linkelist) insertspecific(num, position int) {
	var p *linkelist = node
	var temp = &linkelist{}
	for i := 0; i < position; i++ {
		p = p.next
	}
	temp.number = num
	temp.next = p.next.next
	p.next = temp

}

// Delete specfic position
func (node *linkelist) deletespecific(position int) {
	var p *linkelist = node
	for i := 0; i < position; i++ {
		p = p.next
	}
	p.next = p.next.next
}

// insert starting of linkedlist
func (node *linkelist) insertatbegining(num int) *linkelist {
	var temp = &linkelist{}
	temp.number = num
	temp.next = node.next
	node = temp
	return node
}

// Main function
func main() {
	head := new(linkelist) //creare a head node
	var choice string

	for true {
		fmt.Println("\nEnter your choice")
		fmt.Println("1. Insert value in linked last")
		fmt.Println("2. Display linkedList")
		fmt.Println("3. Deleting from begining ")
		fmt.Println("4. Deleting from last")
		fmt.Println("5. Deleting from specific position")
		fmt.Println("6. Insert at specific position")
		fmt.Println("7. insert at starting")
		fmt.Println("0. Exit")

		fmt.Scanln(&choice)

		switch choice {
		case "1":
			var data string
			fmt.Println("Enter your value for linked list")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			head.insert(num)

		case "2":
			head.display()
		case "3":
			head.deletefirst()
		case "4":
			head.deletelast()
		case "5":
			var sposition string
			fmt.Println("Enter Postion to delete")
			fmt.Scanln(&sposition)
			position, _ := strconv.Atoi(sposition)
			head.deletespecific(position)

		case "6":
			var sposition string
			var data string
			fmt.Println("insert Postion to insert")
			fmt.Scanln(&sposition)
			position, _ := strconv.Atoi(sposition)
			fmt.Println("insert value to add")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			head.insertspecific(num, position)
		case "7":
			var data string
			fmt.Println("Insert value to add")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			head = head.insertatbegining(num)
		default:
			os.Exit(0)
		}
	}
}
