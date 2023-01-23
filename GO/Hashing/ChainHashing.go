package main

import "fmt"

const tableSize = 11

type Node struct {
	key  int
	next *Node
	//value int
}

type HashTable struct {
	items [tableSize]*Node
}

func (ht *HashTable) hash(key int) int {
	return key % tableSize
}

func (ht *HashTable) insert(key int) {
	index := ht.hash(key)
	newNode := &Node{key: key, next: ht.items[index]}
	ht.items[index] = newNode
}

func (ht *HashTable) search(key int) bool {
	index := ht.hash(key)
	node := ht.items[index]

	for node != nil {
		if node.key == key {
			return true
		}

		node = node.next
	}

	return false
}

func main() {
	ht := &HashTable{}

	ht.insert(17)
	ht.insert(23)
	ht.insert(31)
	ht.insert(41)
	ht.insert(53)

	fmt.Println(ht.search(17))
	fmt.Println(ht.search(23))
	fmt.Println(ht.search(31))
	fmt.Println(ht.search(41))
	fmt.Println(ht.search(53))
	fmt.Println(ht.search(67))
}
