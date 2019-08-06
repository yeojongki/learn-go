package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

func createTreeNode(value int) *treeNode {
	return &treeNode{value: value}
}

// 值传递
func (node treeNode) setValueByValue(value int) {
	node.value = value
}

// 引用传递
func (node *treeNode) setValueByRef(value int) {
	if node == nil {
		fmt.Println("ignore nil node.")
		return
	}
	node.value = value
}

// 中序遍历树
func (root *treeNode) traverse() {
	if root == nil {
		return
	}
	root.left.traverse()
	root.print()
	root.right.traverse()
}

func main() {
	root := treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createTreeNode(2)
	fmt.Println(root)

	root.print()
	// nodes := []treeNode{
	// 	{value: 3},
	// 	{},
	// 	{5, nil, &root},
	// }
	// fmt.Println(nodes)

	root.right.left.setValueByValue(1)
	root.right.left.print()

	root.right.left.setValueByRef(1)
	root.right.left.print()

	root.print()
	root.setValueByRef(100)
	fmt.Println()

	var aRoot *treeNode
	aRoot.setValueByRef(100) // will ignore
	aRoot = &root
	aRoot.setValueByRef(500)
	aRoot.print() // 500
	fmt.Println()

	root.traverse()
}
