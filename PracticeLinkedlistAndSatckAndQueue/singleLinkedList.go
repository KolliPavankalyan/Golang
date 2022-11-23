package main

import "fmt"

//create a struct
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

//insert a node at last
func (obj *List) insertAtLast(num int) {
	newNode := node{element: num}
	if obj.len == 0 {
		obj.head = &newNode
		obj.tail = &newNode
	} else {
		obj.tail.next = &newNode

	}
	obj.tail = &newNode
	obj.len++
}

//display the linkedlist
func (obj *List) dispay() {
	temp := obj.head
	for i := 0; i < obj.len; i++ {
		fmt.Printf("%d ->", temp.element)
		temp = temp.next
	}
	fmt.Println()
}

// insert at staring
func (obj *List) insertAtBegining(num int) {
	newNode := node{element: num}
	temp := obj.head
	if obj.len == 0 {
		obj.head = &newNode
		obj.head = &newNode
	} else {
		obj.head = &newNode
		newNode.next = temp
	}
	obj.len++
}

// insert at position
func (obj *List) insertAtPosition(pos, num int) {
	newNode := node{element: num}
	comm := obj.head
	for i := 1; i < pos-1; i++ {
		comm = comm.next
	}
	temp := comm.next
	comm.next = &newNode
	newNode.next = temp
	obj.len++

}

// delete at last
func (obj *List) deleteLast() {
	if obj.len == 0 {
		fmt.Println("No vnodes found")
	}
	obj.tail.next = nil
	obj.tail = nil
	obj.len--
}

// delete at begining
func (obj *List) DeleteFromBeginging() {
	comm := obj.head
	if obj.len == 0 {
		fmt.Println("No vnodes found")
	} else {
		obj.head = comm.next
	}
	obj.len--
}

//find length of linkedlist
func (obj *List) findSize() int {
	return obj.len
}

//serch of node
func (obj *List) serch(val int) bool {
	temp := obj.head
	for i := 0; i < obj.len; i++ {
		if val == temp.element {
			return true
		}
		temp = temp.next
	}
	return false
}

func main() {
	obj := List{}
	obj.insertAtLast(10)
	obj.insertAtLast(20)
	obj.dispay()
	obj.insertAtBegining(30)
	obj.dispay()
	obj.insertAtPosition(1, 50)
	obj.dispay()
	//obj.deleteLast()
	//obj.dispay()
	obj.DeleteFromBeginging()
	obj.dispay()
	fmt.Println(obj.findSize())
	obj.serch(50)
	fmt.Println(obj.serch(12))
	obj.insertAtLast(10)
	obj.insertAtLast(20)
	obj.insertAtLast(30)
	obj.dispay()

}
