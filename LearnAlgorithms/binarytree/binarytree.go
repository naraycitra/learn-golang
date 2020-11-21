package main

import "fmt"

// Node represent binary tree struct
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

var count int

// Insert method that will insert new node,
// it will be inserted to the left if key smaller than current node
// and will be inserted to the right if key higher than current node
func (node *Node) Insert(key int) {
	switch cnode := node.Key; {
	case cnode < key:
		if node.Right != nil {
			node.Right.Insert(key)
		} else {
			node.Right = &Node{Key: key}
		}
	default:
		if node.Left != nil {
			node.Left.Insert(key)
		} else {
			node.Left = &Node{Key: key}
		}
	}
}

// Search method that will search the given value
// and will return true if exist
func (node *Node) Search(key int) bool {
	if node == nil {
		return false
	}
	switch cnode := node.Key; {
	case cnode < key:
		return node.Right.Search(key)
	case cnode > key:
		return node.Left.Search(key)
	default:
		return true
	}
}

func main() {
	root := &Node{Key: 100}
	root.Insert(200)
	root.Insert(50)
	root.Insert(30)
	root.Insert(540)
	root.Insert(330)
	root.Insert(12)
	fmt.Println(root)
	fmt.Println(root.Search(50))
	fmt.Println(root.Search(150))
}
