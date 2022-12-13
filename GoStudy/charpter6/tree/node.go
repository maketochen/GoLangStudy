package main

import "fmt"

func main() {

}

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode) Print() {
	fmt.Print(node.value)
}
