package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

var root = new(Node)

func addNote(t *Node, v int) int {
	if root == nil {
		t := &Node{
			Value: v,
			Next:  nil,
		}
		root = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node already exit:", v)
		return -1
	}

	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}

	return addNote(t.Next, v)
}

func traverseLinked(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list !")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}

	fmt.Println()
}

func lookupNode(t *Node, v int) bool {
	if root == nil {
		t = &Node{v, nil}
		root = t
		return false
	}
	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty list !")
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}

	return i
}

func MainLinkedList() {
	fmt.Println(root)
	root = nil
	traverseLinked(root)
	addNote(root, 1)
	addNote(root, -1)
	traverseLinked(root)
	addNote(root, 10)
	addNote(root, 5)
	addNote(root, 50)
	addNote(root, 4)
	addNote(root, 4)

	traverseLinked(root)

	if lookupNode(root, 50) {
		fmt.Println("Node exits !")

	} else {
		fmt.Println("Node does not exit")
	}

	if lookupNode(root, 100) {
		fmt.Printf("Node exits")
	} else {
		fmt.Println("Node does not exit !")
	}

	fmt.Println(size(root))

}
