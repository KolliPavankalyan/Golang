package main

import (
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	value int
	next  *Node
}

type CircularLinked struct {
	head *Node
	tail *Node
	len  int
}

func (c *CircularLinked) addLast(val int) {
	newNode := &Node{value: val}
	if c.head == nil {
		newNode.next = newNode
		c.head = newNode

	} else {
		newNode.next = c.tail.next
		c.tail.next = newNode
		
	}
	c.tail = newNode
	c.len++
}

func (c *CircularLinked) addFirst(val int) {
	newNode := &Node{value: val}
	if c.len == 0 {
		newNode.next = newNode
		c.head = newNode
	} else {
		c.tail.next = newNode
		newNode.next = c.head
		c.head = newNode
	}
	c.len++
}

func (c *CircularLinked) insertAny(val, pos int) {
	newNode := &Node{value: val}
	ptr := c.head
	for i := 1; i < pos-1; i++ {
		ptr = ptr.next
	}
	newNode.next = ptr.next
	ptr.next = newNode
	c.len++
}

func (c *CircularLinked) display() {
	p := c.head
	for i := 0; i < c.len; i++ {
		fmt.Printf("%d -> ", p.value)
		p = p.next
	}
}

func (c *CircularLinked) delLast() {
	ptr := c.head
	for i := 0; i < c.len-1; i++ {
		ptr = ptr.next
	}
	ptr.next = c.head
	c.len--
}
func (c *CircularLinked) delFirst() {
	c.tail.next = c.head.next
	c.head = c.head.next
	c.len--
}

func (c *CircularLinked) delAny(pos int) {
	ptr := c.head
	for i := 1; i < pos-1; i++ {
		ptr = ptr.next
	}
	ptr.next = ptr.next.next
	c.len--
}

func main() {
	c := CircularLinked{}

	var choice string

	for true {
		fmt.Println("please enter choice")
		fmt.Println("1. Insert At end")
		fmt.Println("2. Display linkelist")
		fmt.Println("3. Insert values at starting")
		fmt.Println("4. Insert at any point")
		fmt.Println("5. Delete first at linkedlist")
		fmt.Println("6. Delete at end of the linkedlist")
		fmt.Println("7. Delete at any point")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			var data string
			fmt.Println("enter value")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			c.addLast(num)
		case "2":
			c.display()
		case "3":
			var data string
			fmt.Println("enter value")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			c.addFirst(num)
		case "4":
			var sposition string
			var data string
			fmt.Println("insert Postion to insert")
			fmt.Scanln(&sposition)
			position, _ := strconv.Atoi(sposition)
			fmt.Println("insert value to add")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			c.insertAny(num, position)
		case "5":
			c.delFirst()
		case "6":
			c.delLast()
		case "7":
			var sposition string
			fmt.Println("insert Postion to delete")
			fmt.Scanln(&sposition)
			position, _ := strconv.Atoi(sposition)
			c.delAny(position)

		default:
			os.Exit(0)
		}
	}
}
