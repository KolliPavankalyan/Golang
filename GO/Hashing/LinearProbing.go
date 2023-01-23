package main

import "fmt"

const tableSize = 11

type HashTable struct {
	items [tableSize]int
}

func (ht *HashTable) hashCode(key int) int {
	return key % tableSize
}

func (ht *HashTable) insert(key int) {
	index := ht.hashCode(key)

	for ht.items[index] != 0 {
		index = (index + 1) % tableSize
	}

	ht.items[index] = key
}

func (ht *HashTable) search(key int) bool {
	index := ht.hashCode(key)

	for ht.items[index] != key {
		index = (index + 1) % tableSize

		if ht.items[index] == 0 {
			return false
		}
	}

	return true
}

func main() {
	ht := &HashTable{}

	ht.insert(17)
	ht.insert(23)
	ht.insert(31)
	ht.insert(41)
	ht.insert(53)

	fmt.Println(ht.search(17)) //true
	fmt.Println(ht.search(23)) //true
	fmt.Println(ht.search(31)) //true
	fmt.Println(ht.search(41)) //true
	fmt.Println(ht.search(53)) //true
	fmt.Println(ht.search(67)) //false
}
