package main

import "fmt"

//create a node class
type Node struct {
	element int
	next    *Node
}

//create List class
type List struct {
	head *Node
	len  int
	tail *Node
}

//printing each node value
func (obj *List) Display() {
	comm := obj.head
	for i := 0; i < obj.len; i++ {
		fmt.Printf("%v==>", comm.element)
		comm = comm.next
	}
	fmt.Println()
}

func (obj *List) count() int {
	return obj.len
}

//adding nodes to the list
func (obj *List) adding(val int) {
	newnode := Node{element: val}
	if obj.len == 0 {
		obj.head = &newnode
	} else {
		obj.tail.next = &newnode
	}
	obj.tail = &newnode
	obj.len++

}

//searching a value node present or not
func (obj *List) search(val int) bool {
	comm := obj.head
	for i := 0; i < obj.len; i++ {
		check := comm.element
		if check == val {
			return true
		}
		comm = comm.next
	}
	return false

}

//append node at the beginning of the list
func (obj *List) appendstart(val int) {
	newnode := Node{element: val}
	temp := obj.head
	if obj.len == 0 {
		obj.head = &newnode
		obj.tail = &newnode
	} else {
		obj.head = &newnode
		newnode.next = temp
	}
	obj.len++
}

//insert node at any position
func (obj *List) insertat(pos, val int) {
	newnode := Node{element: val}
	comm := obj.head
	for i := 1; i < pos-1; i++ {
		comm = comm.next
	}
	temp := comm.next
	comm.next = &newnode
	newnode.next = temp
	obj.len++

}

//delete a node at beginning
func (obj *List) deletebegg() {
	if obj.len == 0 {
		fmt.Println("no nodes found")
		return
	}
	comm := obj.head
	obj.head = comm.next
	obj.len--
}

//delete a node at the end
func (obj *List) deleteendd() {
	if obj.len == 0 {
		fmt.Println("no nodes are found")
		return
	}
	obj.tail.next = nil
	obj.tail = nil
	obj.len--

}

//delete at pos
func (obj *List) DeleteAtPos(pos int) {
	comm := obj.head
	for i := 1; i < pos-1; i++ {
		comm = comm.next
	}
	comm.next = comm.next.next
}

func main() {
	obj := List{}
	obj.adding(10)
	obj.adding(20)
	obj.adding(30)
	obj.adding(40)

	fmt.Println(obj.count())

	obj.Display()

	fmt.Println(obj.search(10))

	fmt.Println(obj.search(1340))

	status3 := obj.search(20)
	fmt.Println(status3)

	obj.appendstart(1000)
	obj.Display()
	obj.adding(2000)
	obj.Display()

	obj.insertat(4, 3000)
	obj.Display()

	obj.deletebegg()
	obj.Display()

	obj.deleteendd()
	obj.Display()

}
