package main

import "fmt"

type Node struct {
	prev    *Node
	element int
	next    *Node
}

type List struct {
	head *Node
	tail *Node
	len  int
}

// print each node value in forward way
func (obj *List) DisplayForward() {
	comm := obj.head
	for i := 0; i < obj.len; i++ {
		fmt.Printf("%v==", comm.element)
		comm = comm.next
	}
	fmt.Println()
}

//print each node value in reverse order
func (obj *List) DisplayReverse() {
	comm := obj.tail
	for i := 0; i < obj.len; i++ {
		fmt.Printf("%v==", comm.element)
		comm = comm.prev
	}
	fmt.Println()
}

// creating nodes and append to list
func (obj *List) InsertLast(val int) {
	newnode := Node{element: val}
	if obj.len == 0 {
		obj.head = &newnode
		obj.tail = &newnode
	} else {
		obj.tail.next = &newnode
		newnode.prev = obj.tail
	}
	obj.tail = &newnode
	obj.len++
}

//create a node and append at beginning of list
func (obj *List) insertbegg(val int) {
	newnode := Node{element: val}
	if obj.len == 0 {
		obj.head = &newnode
		obj.tail = &newnode
	} else {
		temp := obj.head
		newnode.next = temp
		temp.prev = &newnode
		obj.head = &newnode
	}
	obj.len++
}

//insert a node at given position in the list
func (obj *List) insertAtPos(pos, val int) {
	newnode := Node{element: val}
	comm := obj.head
	for i := 1; i < pos-1; i++ {
		comm = comm.next
	}
	newnode.prev = comm
	newnode.next = comm.next
	comm.next = &newnode
	comm.next.prev = &newnode
	obj.len++
}

//delete node at the beginning
func (obj *List) deletebegg() {
	comm := obj.head
	obj.head = comm.next
	obj.head.prev = nil
	obj.len--

}

//delete node at the end
func (obj *List) deleteend() {
	comm := obj.tail
	comm.prev.next = nil
	obj.tail = comm.prev.next
	obj.len--
}

//delete a node at given value
func (obj *List) deleteat(pos int) int {
	comm := obj.head
	for i := 1; i < pos-1; i++ {
		comm = comm.next
	}
	value := comm.next.element
	comm.next = comm.next.next
	comm.next.prev = comm
	obj.len--
	return value
}

func main() {
	obj := List{}
	obj.InsertLast(10)
	obj.InsertLast(20)
	obj.InsertLast(30)
	obj.InsertLast(40)

	obj.DisplayForward()

	obj.DisplayReverse()

	obj.insertbegg(50)

	obj.DisplayForward()

	obj.InsertLast(60)

	obj.DisplayForward()

	obj.insertAtPos(3, 70)

	obj.DisplayForward()

	obj.deletebegg()

	obj.DisplayForward()

	obj.deleteend()

	obj.DisplayForward()

	fmt.Println(obj.deleteat(3))

	obj.DisplayForward()
}
