package main

import (
	"ccmouse-go/C6/tree"
	"fmt"
)

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(4)

	root.Traverse()

	TraverseCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		TraverseCount++
	})
	fmt.Println("TraverseCount: ", TraverseCount)

	myNode := MyNode{&root}
	myNode.postOrder()
	fmt.Println()
}

//使用组合的方式扩展原有类型
type MyNode struct {
	node *tree.Node
}

//倒序打印树
func (myNode *MyNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	leftNode := MyNode{myNode.node.Left}
	rightNode := MyNode{myNode.node.Right}
	leftNode.postOrder()
	rightNode.postOrder()
	myNode.node.Print()
}
