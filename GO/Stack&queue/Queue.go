package main

import (
	"fmt"
	"os"
	"strconv"
)

// creating a struct
type node struct {
	number int
	next   *node
}

// creating a class queue
type Queue struct {
	len  int
	font *node
	rear *node
}

// insert a value to Queue
func (head *Queue) Enqueue(num int) {
	var temp = &node{}
	temp.number = num
	temp.next = nil
	if head.font == nil && head.rear == nil {
		head.font = temp
		head.rear = temp
	} else {
		head.rear.next = temp
		head.rear = temp
	}
	head.len++
}

// delete a value from queue
func (head *Queue) Dequeue() {
	if head.rear == head.font {
		head.rear = nil
		head.font = nil
	} else {
		head.font = head.font.next
	}
	head.len--
}

// display value of first node
func (head *Queue) Front() {
	fmt.Println("Front node is ", head.font.number)
}

// find length of queue
func (head *Queue) Size() {
	fmt.Println("Size of Queue is ", head.len)
}

// display the entire queue
func (head *Queue) Display() {
	var p *node = head.font
	for p != nil {
		fmt.Printf("%d <-", p.number)
		p = p.next
	}
}

func main() {
	var myqueue = &Queue{}
	var choice string
	for true {
		fmt.Println("Enter your choice")
		fmt.Println("1. Enqueue")
		fmt.Println("2. Dequeue")
		fmt.Println("3. Front")
		fmt.Println("4. size")
		fmt.Println("5. Display/travese")
		fmt.Println("6. Exit")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			var data string
			fmt.Println("Enter your value for linked list")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			myqueue.Enqueue(num)
		case "2":
			myqueue.Dequeue()
		case "3":
			myqueue.Front()
		case "4":
			myqueue.Size()
		case "5":
			myqueue.Display()
		default:
			os.Exit(0)

		}
	}

}
