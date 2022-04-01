package main

import (
	"fmt"
	"github.com/yb19890724/leetcode-go/src/structure"
)

func main() {

	tree := &structure.BinaryTree{}

	tree.Add(1)
	tree.Add(2)
	tree.Add(3)
	tree.Add(4)
	tree.Add(5)
	tree.Add(6)


	fmt.Println(tree.MiddleFor())
}
