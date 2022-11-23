package main

import "fmt"

// create a node
type node struct {
	element int
	next    *node
}

//create a class
type List struct {
	head *node
	len  int
	tail *node
}

// insert a value at end
func (obj *List) instert(val int) {
	newNode := node{element: val}
	if obj.len == 0 {
		obj.head = &newNode
		obj.tail = &newNode
	} else {
		obj.tail.next = &newNode
		obj.tail = &newNode
	}
	obj.len++
}

// display the linkedlist
func (obj *List) dispay() {
	comm := obj.head
	for i := 0; i < obj.len; i++ {
		fmt.Printf("%d ->", comm.element)
		comm = comm.next
	}
	fmt.Println()
}

//insert At Begining
func (obj *List) insertBeg(val int) {
	newNode := node{element: val}
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

//Insert At Position
func (obj *List) insertAtPos(pos, val int) {
	comm := obj.head
	newNode := node{element: val}
	for i := 1; i < pos-1; i++ {
		comm = comm.next
	}
	temp := comm.next
	comm.next = &newNode
	newNode.next = temp
	obj.len++
}

//Delete at Last
func (obj *List) deleteLast() {
	if obj.len == 0 {
		fmt.Println("no nodes found")
	}
	obj.tail.next = nil
	obj.tail = nil
	obj.len--
}

//Delete at beginging
func (obj *List) deleteBeg() {
	if obj.len == 0 {
		fmt.Println("no nodes found")
	}
	obj.head = obj.head.next
	obj.len--
}

//delete at position
func (obj *List) deleteAtPos(pos int) {
	comm := obj.head
	for i := 0; i < pos-1; i++ {
		comm = comm.next
	}
	comm.next = comm.next.next
	obj.len--
}

func main() {
	obj := List{}
	obj.instert(10)
	obj.dispay()
	obj.insertBeg(20)
	obj.insertBeg(50)
	obj.dispay()
	obj.insertAtPos(1, 30)
	obj.dispay()
	obj.deleteLast()
	obj.dispay()
	obj.deleteBeg()
	obj.dispay()
	obj.insertBeg(500)
	obj.dispay()
	obj.deleteAtPos(2)
	obj.dispay()
}
