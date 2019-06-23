package main

import "fmt"

//noinspection GoDuplicate
type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Add(value int) {
	if t.root == nil {
		t.root = &Node{value: value}
		return
	}

	t.addToNode(t.root, value)
}

func (t *Tree) addToNode(node *Node, value int) {
	if value < node.value {
		if node.left == nil {
			node.left = &Node{value: value}
		} else {
			t.addToNode(node.left, value)
		}
		return
	}

	if node.right == nil {
		node.right = &Node{value: value}
	} else {
		t.addToNode(node.right, value)
	}
}

func (t *Tree) Exists(v int) bool {
	if t.root != nil {
		return t.root.exists(v)
	}

	return false
}

func (n *Node) exists(v int) bool {
	if n == nil {
		return false
	}

	if n.value == v {
		return true
	}

	if v <= n.value {
		return n.left.exists(v)
	} else {
		return n.right.exists(v)
	}
}

func main() {
	t := Tree{}
	t.Add(10)
	t.Add(5)
	t.Add(20)
	t.Add(18)
	t.Add(76)
	fmt.Println(t.Exists(77))
}
