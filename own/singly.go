package main

import "fmt"

type node struct {
	element int
	next    *node
}

type list struct {
	head *node
	len  int
	tail *node
}

func (obj *list) insertAtEnd(val int) {
	newNode := node{element: val}
	if obj.len == 0 {
		obj.head = &newNode
		obj.tail = &newNode
	}
	obj.tail.next = &newNode
	obj.tail = &newNode
	obj.len++
}

func (obj *list) display() {
	comm := obj.head
	for i := 0; i < obj.len; i++ {
		fmt.Printf("%d ->", comm.element)
		comm = comm.next
	}
	fmt.Println()
}

func (obj *list) insertAtBeg(val int) {
	newnode := node{element: val}

	if obj.len == 0 {
		obj.head = &newnode
		obj.tail = &newnode
	}
	temp := obj.head
	newnode.next = temp
	obj.head = &newnode
	obj.len++
}

func (obj *list) InsertAtPos(pos, val int) {
	newnode := node{element: val}
	comm := obj.head
	for i := 1; i < pos-1; i++ {
		comm = comm.next
	}

	temp := comm.next
	comm.next = &newnode
	newnode.next = temp
	obj.len++

}

func main() {
	obj := list{}
	obj.insertAtEnd(10)
	obj.insertAtEnd(20)
	obj.display()
	obj.insertAtBeg(30)
	obj.display()
	obj.InsertAtPos(2, 40)
	obj.display()
}
