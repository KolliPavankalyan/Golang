package main

import (
	"fmt"
	"os"
	"strconv"
)

type node struct {
	number int
	prev   *node
	next   *node
}

type doublylinkedlist struct {
	len  int
	head *node
	tail *node
}

// insert value at last
func (dll *doublylinkedlist) InsertAtlast(num int) {
	var temp = &node{}
	temp.number = num
	temp.prev = nil
	temp.next = nil
	if dll.head == nil && dll.tail == nil {
		dll.head = temp
		dll.tail = temp
	} else {
		temp.prev = dll.tail
		dll.tail.next = temp
		dll.tail = temp
	}
	dll.len++
}

// Display/Traverse
func (dll *doublylinkedlist) Display() {
	var p *node = dll.head
	for p != nil {
		fmt.Printf("%d ->", p.number)
		p = p.next
	}
}

// Display Reverse order
func (dll *doublylinkedlist) DisplayReverse() {
	var p *node = dll.tail
	for p != nil {
		fmt.Printf("%d ->", p.number)
		p = p.prev
	}
}

// Insert at begining
func (dll *doublylinkedlist) InsertAtBegining(num int) {
	var temp = &node{}
	temp.number = num
	temp.prev = nil
	temp.next = nil
	if dll.head == nil && dll.tail == nil {
		dll.head = temp
		dll.tail = temp
	} else {
		temp.next = dll.head
		dll.head.prev = temp
		dll.head = temp
	}
	dll.len++
}

// delete from Begining
func (dll *doublylinkedlist) DeleteFromBeginging() {
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.head = dll.head.next
		dll.head.prev = nil
	}
	dll.len--
}

// delete form end of linkedlist
func (dll *doublylinkedlist) DeleteFromEnd() {
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.tail = dll.tail.prev
		dll.tail.next = nil
	}
	dll.len--
}

// Delete Specific index
func (dll *doublylinkedlist) DeleteFromSpecific(position int) {
	if dll.len >= position {
		var p *node = dll.head
		for i := 0; i < position-1; i++ {
			p = p.next
		}
	}
}

func main() {
	var dll = &doublylinkedlist{}
	var choice string
	for true {
		fmt.Println("Enter your choice")
		fmt.Println("1. Insert node at end")
		fmt.Println("2. Traverse/Display")
		fmt.Println("3. Travarse /Display in reverse order")
		fmt.Println("4. Insert at Begging")
		fmt.Println("5. Delete from begining")
		fmt.Println("6. Delete from end")
		fmt.Println("7. Delete from specific position")
		fmt.Println("8. Insert at specific position")
		fmt.Println("0. Exit")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			var data string
			fmt.Println("please enter number")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			dll.InsertAtlast(num)
		case "2":
			dll.Display()
		case "3":
			dll.DisplayReverse()
		case "4":
			var data string
			fmt.Println("enter number")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			dll.InsertAtBegining(num)
		case "5":
			dll.DeleteFromBeginging()
		case "6":
			dll.DeleteFromEnd()
		case "7":
			var sposition string
			fmt.Println("Enter the position")
			fmt.Scanln(&sposition)
			position, _ := strconv.Atoi(sposition)
			dll.DeleteFromSpecific(position)

		default:
			os.Exit(0)
		}
	}
}
