package main

import "fmt"

// create a node
type Node struct {
	element int
	next    *Node
}

//create a class
type List struct {
	head *Node
	len  int
	tail *Node
}

// insert a node at end
func (obj *List) InsertAtEnd(val int) {
	newNode := Node{element: val}

	if obj.len == 0 {
		obj.head = &newNode
		obj.tail = &newNode
	} else {
		obj.tail.next = &newNode
	}
	obj.tail = &newNode
	obj.len++
}

//display linked list
func (obj *List) Display() {
	comm := obj.head
	for i := 0; i < obj.len; i++ {
		fmt.Printf("%d ->", comm.element)
		comm = comm.next
	}
	fmt.Println()
}

// instert at starting of linkedlist
func (obj *List) AddAtBegining(val int) {
	newNode := Node{element: val}
	temp := obj.head
	if obj.len == 0 {
		obj.head = &newNode
		obj.tail = &newNode
	} else {
		obj.head = &newNode
		newNode.next = temp
	}
	obj.len++
}

// search element in linkedlist
func (obj *List) searchNode(val int) bool {
	comm := obj.head
	for i := 0; i < obj.len; i++ {
		check := comm.element
		if check == val {
			return true
		} else {
			comm = comm.next
		}
	}
	return false
}

// instert at position
func (obj *List) InsertAtPostion(pos, value int) {
	newNode := Node{element: value}
	comm := obj.head

	for i := 1; i < pos-1; i++ {
		comm = comm.next
	}
	temp := comm.next
	comm.next = &newNode
	newNode.next = temp
	obj.len++
}

// delete at beginging in linkedlist
func (obj *List) deleteAtBegining() {
	comm := obj.head
	if obj.len == 0 {
		fmt.Println("no nodes are found")
	}
	obj.head = comm.next
	obj.len--

}

// delete at ending in linkedlist
func (obj *List) deleteAtEnd() {
	if obj.len == 0 {
		fmt.Println("no nodes found")
	} else {
		obj.tail.next = nil
		obj.tail = nil
	}
	obj.len--
}

// len of linkedlist
func (obj *List) FindLength() {
	fmt.Println(obj.len)
}
func main() {
	obj := List{}
	obj.InsertAtEnd(20)
	obj.InsertAtEnd(30)
	obj.Display()
	obj.InsertAtEnd(35)
	obj.Display()
	obj.AddAtBegining(10)
	obj.Display()
	obj.InsertAtPostion(2, 50)
	obj.Display()
	obj.deleteAtBegining()
	obj.Display()
	obj.deleteAtEnd()
	obj.Display()
	obj.FindLength()
}
