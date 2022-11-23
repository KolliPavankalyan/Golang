package main

import (
	"fmt"
	"os"
	"strconv"
)

type stack struct {
	element int
	next    *stack
}

func (Node *stack) push(val int) *stack {
	var temp = stack{element: val}
	temp.next = Node
	Node = &temp

	return Node

}

func (Node *stack) pop() *stack {
	Node = Node.next
	if Node == nil {
		fmt.Println("Stack is empty")
	}
	return Node
}

func (Node *stack) peek() {
	fmt.Println(Node.element)
}

func (Node *stack) display() {
	var p *stack = Node
	for p.next != nil {
		fmt.Printf("| %d |", p.element)
		fmt.Println("\n____________")
		p = p.next
	}
}
func main() {
	head := new(stack)
	var choice string
	for true {
		fmt.Println("Enter your choice")
		fmt.Println("1. Push value in stack")
		fmt.Println("2. Pop value from stack")
		fmt.Println("3. peep value from stack")
		fmt.Println("4. Display stack")
		fmt.Println("5. Exit")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			var data string
			fmt.Println("Enter your value for linked list")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)

			head = head.push(num)
		case "2":
			head = head.pop()
		case "3":
			head.peek()
		case "4":
			head.display()
		default:
			os.Exit(0)

		}
	}
}
