package main

import "fmt"

//create a node
type Node struct {
	element int
	next    *Node
}

//create a class
type List struct {
	head *Node
	tail *Node
	len  int
}

// count of the list
func (obj *List) Length() int {
	return obj.len
}

//display each node in the list
func (obj *List) Display() {
	comm := obj.head
	for i := 0; i < obj.len; i++ {
		fmt.Printf("%v==>", comm.element)
		comm = comm.next
	}
	fmt.Println()
}

// create a node and append at last in list
func (obj *List) InsertAtLast(val int) {
	newnode := Node{element: val}

	if obj.len == 0 {
		obj.head = &newnode
		obj.tail = &newnode
	} else {
		newnode.next = obj.tail.next
		obj.tail.next = &newnode
	}
	obj.tail = &newnode
	obj.len++
}

//append a node at beginning
func (obj *List) InsertAtBegining(val int) {
	newnode := Node{element: val}
	if obj.len == 0 {
		obj.head = &newnode
		obj.tail = &newnode
		newnode.next = &newnode
	} else {
		obj.tail.next = &newnode
		newnode.next = obj.head
		obj.head = &newnode
	}
	obj.len++

}

//insert a node at position
func (obj *List) InsertAtPosition(pos, val int) {
	newnode := Node{element: val}
	comm := obj.head
	for i := 1; i < pos-1; i++ {
		comm = comm.next
	}
	newnode.next = comm.next
	comm.next = &newnode
	obj.len++
}

//delete node at starting
func (obj *List) DeleteAtBegining() {
	obj.head = obj.head.next
	obj.tail.next = obj.head
	obj.len--
}

//delete node at end of the list
func (obj *List) DeleteEnd() {
	comm := obj.head
	for i := 1; i < obj.len-2; i++ {
		comm = comm.next
	}
	add := comm.next
	add.next = obj.head
	obj.len--
}

//delete a node in between nodes
func (obj *List) DeleteAtPosition(pos int) {
	comm := obj.head
	for i := 1; i < pos-1; i++ {
		comm = comm.next
	}
	comm.next = comm.next.next
	obj.len--

}
func main() {
	obj := List{}
	obj.InsertAtLast(10)
	obj.InsertAtLast(20)
	obj.InsertAtLast(30)
	obj.InsertAtLast(40)
	fmt.Println(obj.Length())
	obj.Display()

	obj.InsertAtBegining(1000)
	obj.Display()

	obj.InsertAtLast(3000) //adding node at end
	obj.Display()

	obj.InsertAtPosition(4, 2000)
	obj.Display()

	obj.DeleteAtBegining()
	obj.Display()

	obj.DeleteEnd()
	obj.Display()

	obj.DeleteAtPosition(3)
	obj.Display()

}
