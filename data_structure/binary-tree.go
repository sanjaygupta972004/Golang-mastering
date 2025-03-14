package main

import (
	"fmt"
	"math/rand"

	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

func create(v int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*v; i++ {
		temp := rand.Intn(v * 2)
		t = t.insert(temp)
	}
	return t
}

func (t *Tree) insert(v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}

	if v == t.Value {
		return t
	}

	if v < t.Value {
		t.Left = t.Left.insert(v)
		return t
	}

	t.Right = t.Right.insert(v)
	return t

}

func (t *Tree) Search(k int) bool {

	if t == nil {
		return false
	}
	if k > t.Value {
		return t.Right.Search(k)

	} else if k < t.Value {
		return t.Left.Search(k)
	}

	return true
}

func BinaryTree() {
	tree := create(10)
	fmt.Println("The value of the root of the tree is", tree.Value)
	traverse(tree)
	tree.insert(20)
	traverse(tree)
	fmt.Println(tree.Search(2))
	fmt.Println("The value of the root of the tree is ", tree.Value)
}
