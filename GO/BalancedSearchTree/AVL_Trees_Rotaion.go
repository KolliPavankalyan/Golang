package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leftRotate(root *TreeNode) *TreeNode {
	newRoot := root.Left
	root.Left = newRoot.Right
	newRoot.Right = root
	return newRoot
}

func rightRotate(root *TreeNode) *TreeNode {
	newRoot := root.Right
	root.Right = newRoot.Left
	newRoot.Left = root
	return newRoot
}

func leftRightRotate(root *TreeNode) *TreeNode {
	root.Left = rightRotate(root.Left)
	return leftRotate(root)
}

func rightLeftRotate(root *TreeNode) *TreeNode {
	root.Right = leftRotate(root.Right)
	return rightRotate(root)
}

func main() {
	// Create a binary tree for testing
	root := &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 15}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 7}
	root.Right.Left = &TreeNode{Val: 12}
	root.Right.Right = &TreeNode{Val: 20}

	fmt.Println("Original tree:")
	printTree(root)

	fmt.Println("\nAfter left-left rotation:")
	root = leftRotate(root)
	printTree(root)

	fmt.Println("\nAfter right-right rotation:")
	root = rightRotate(root)
	printTree(root)

	fmt.Println("\nAfter left-right rotation:")
	root = leftRightRotate(root)
	printTree(root)

	fmt.Println("\nAfter right-left rotation:")
	//root = rightLeftRotate(root)
	printTree(root)
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	printTree(root.Left)
	printTree(root.Right)
}
