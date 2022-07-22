package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func newNode(data int) *Node {
	return &Node{data: data, left: nil, right: nil}
}

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return new(Tree)
}

func (t *Tree) Insert(data int) {
	t.root = insert(t.root, data)
}

func (t *Tree) Remove(data int) {
	t.root = remove(t.root, data)
}

func insert(node *Node, data int) *Node {
	if node == nil {
		return newNode(data)
	}

	if node.data > data {
		node.left = insert(node.left, data)
	} else {
		node.right = insert(node.right, data)
	}

	return node
}

func searchMin(node *Node) int {
	if node.left == nil {
		return node.data
	}
	return searchMin(node.left)
}

func deleteMin(node *Node) *Node {
	if node.left == nil {
		return node.right
	}
	node.left = deleteMin(node.left)
	return node
}

func remove(node *Node, data int) *Node {
	if node == nil {
		return nil
	}

	switch {
	case node.data > data:
		node.left = remove(node.left, data)
	case node.data < data:
		node.right = remove(node.right, data)
	case node.data == data:
		switch {
		case node.left == nil:
			return node.right
		case node.right == nil:
			return node.left
		default:
			node.data = searchMin(node.right)
			node.right = deleteMin(node.right)
		}
	}

	return node
}

func main() {
	t := NewTree()
	t.Insert(10)
	t.Insert(20)
	t.Insert(5)
	fmt.Printf("t: %v\n", t)

	t.Remove(10)
	fmt.Printf("t: %v\n", t)

}
