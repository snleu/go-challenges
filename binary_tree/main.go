package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) Add(value int) {
	if n == nil {
		return
	} else if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value, left: nil, right: nil}
		} else {
			n.left.Add(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value, left: nil, right: nil}
		} else {
			n.right.Add(value)
		}
	}
}

func print(node *Node, ns int, name string) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%v:%v\n", name, node.value)
	print(node.left, ns+2, "L")
	print(node.right, ns+2, "R")
}

func main() {
	node := &Node{value: 5, left: nil, right: nil}
	node.Add(3)
	node.Add(9)
	node.Add(4)
	node.Add(10)
	node.Add(7)
	print(node, 0, "Root")
}
