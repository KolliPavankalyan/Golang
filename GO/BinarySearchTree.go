package main

import "fmt"

type Node struct {
	element int
	left    *Node
	right   *Node
}

type BinaryTree struct {
	root *Node
}

// Create a Root
func (obj *BinaryTree) MakeTree(val int) {

	if obj.root == nil {
		obj.root = &Node{element: val, left: nil, right: nil}
		return
	}

	obj.root.insert(val)
}

// Create a Binary Tree
func (curr_node_add *Node) insert(data int) {

	if data <= curr_node_add.element {

		if curr_node_add.left == nil {
			curr_node_add.left = &Node{element: data, left: nil, right: nil}

		} else {
			curr_node_add.left.insert(data)
		}

	} else {

		if curr_node_add.right == nil {
			curr_node_add.right = &Node{element: data, left: nil, right: nil}

		} else {
			curr_node_add.right.insert(data)
		}
	}

}

// Print PreOrder
func (obj *BinaryTree) Preorder(root *Node) {

	if root != nil {
		fmt.Printf("%d ", root.element)
		obj.Preorder(root.left)
		obj.Preorder(root.right)
	}
}

//Print Inorder
func (b *BinaryTree) Inorder(troot *Node) {
	if troot != nil {
		b.Preorder(troot.left)
		fmt.Printf("%d ", troot.element)
		b.Preorder(troot.right)
	}
}

// Print PostOrder
func (b *BinaryTree) Postorder(troot *Node) {
	if troot != nil {
		b.Preorder(troot.left)
		b.Preorder(troot.right)
		fmt.Printf("%d ", troot.element)
	}
}

//LevelOrder
func (b *BinaryTree) levelorder(troot *Node) {
	Q := []*Node{}
	t := troot
	fmt.Print(troot.element, " ")
	Q = append(Q, t)
	for len(Q) > 0 {
		t = Q[0]
		Q = Q[1:]
		if t.left != nil {
			fmt.Print(t.left.element, " ")
			Q = append(Q, t.left)
		}
		if t.right != nil {
			fmt.Print(t.right.element, " ")
			Q = append(Q, t.right)
		}
	}
}

//Height of the tree

func (b *BinaryTree) Height(troot *Node) int {
	if troot != nil {
		x := b.Height(troot.left)
		y := b.Height(troot.right)
		if x > y {
			return x + 1
		} else {
			return y + 1
		}
	}
	return 0
}

//Search a element
func (bt *Node) find(val int) bool {
	troot := bt
	for troot != nil {
		if troot.element == val {
			return true
		} else if troot.element < val {
			troot = troot.right
		} else if troot.element > val {
			troot = troot.left
		}
	}
	return false

}

//main Function
func main() {
	obj := BinaryTree{}
	sli := []int{100, 50, 40, 60, 200, 250, 120}
	for _, v := range sli {
		obj.MakeTree(v)
	}

	fmt.Println("Preorder ---->")
	obj.Preorder(obj.root)

	fmt.Println("\nInorder ---->")
	obj.Inorder(obj.root)

	fmt.Println("\nPostorder ---->")
	obj.Postorder(obj.root)

	fmt.Println("\nLevelOrder ---->")
	obj.levelorder(obj.root)

	fmt.Println("\nHeight of Tree ---->")
	fmt.Println(obj.Height(obj.root))

	fmt.Println(obj.root.find(2500))

}
