package tree

func (root *Node) Traverse() {
	if root == nil {
		return
	}
	root.Left.Traverse()
	root.Print()
	root.Right.Traverse()
}
