package main

import (
	"fmt"
	"math/rand"
	"time"
)

type node struct {
	left  *node
	right *node
	data  int
}

func (n *node) insert(d int) {
	if d <= n.data {
		if n.left == nil {
			n.left = &node{data: d}
		} else {
			n.left.insert(d)
		}
	} else {
		if n.right == nil {
			n.right = &node{data: d}
		} else {
			n.right.insert(d)
		}
	}
}

func print(n *node) {
	if n == nil {
		return
	}

	print(n.left)
	fmt.Printf("%d ", n.data)
	print(n.right)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	size := 10
	root := &node{data: rand.Intn(999)}
	for i := 1; i < size; i++ {
		root.insert(rand.Intn(999))
	}

	print(root)
}
