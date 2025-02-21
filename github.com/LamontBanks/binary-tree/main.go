package main

import (
	"fmt"
	"math/rand"
)

func main() {

	root := TreeNode{
		val:   10,
		left:  nil,
		right: nil,
	}

	numVals := 25
	for i := 0; i < numVals; i++ {
		root.insertNode(rand.Intn(100))
	}

	fmt.Println(root.inOrderPrint())
}

// ---

type TreeNode struct {
	left, right *TreeNode
	val         int
}

func (node *TreeNode) insertNode(newValue int) {
	if newValue <= node.val {
		if node.left != nil {
			node.left.insertNode(newValue)
		} else {
			node.left = &TreeNode{
				val:   newValue,
				left:  nil,
				right: nil,
			}
		}
	}

	if newValue > node.val {
		if node.right != nil {
			node.right.insertNode(newValue)
		} else {
			node.right = &TreeNode{
				val:   newValue,
				left:  nil,
				right: nil,
			}
		}
	}
}

func (node TreeNode) inOrderPrint() string {
	var str string

	if node.left != nil {
		str += node.left.inOrderPrint()
	}

	str += fmt.Sprintf("%v|", node.val)

	if node.right != nil {
		str += node.right.inOrderPrint()
	}

	return str
}
