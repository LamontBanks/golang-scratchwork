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
	randMin := rand.Intn(100)
	randMax := rand.Intn(100) + randMin

	fmt.Printf("Values between %v and %v:\n%v", randMin, randMax, root.getRange(randMin, randMax))
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

// Get values within the min, max range
func (node TreeNode) getRange(minVal, maxVal int) []int {
	var values []int

	if node.val >= minVal {
		if node.left != nil {
			values = append(values, node.left.getRange(minVal, maxVal)...)
		}
	}

	if node.val >= minVal && node.val <= maxVal {
		values = append(values, node.val)
	}

	if node.val <= maxVal {
		if node.right != nil {
			values = append(values, node.right.getRange(minVal, maxVal)...)
		}
	}

	return values
}
