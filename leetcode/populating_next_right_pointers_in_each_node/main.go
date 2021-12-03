package populating_next_right_pointers_in_each_node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}
	recursiveLinkPairs(root.Left, root.Right)
	return root
}

func recursiveLinkPairs(leftNode, rightNode *Node) {
	if leftNode.Next != nil {
		return
	}

	leftNode.Next = rightNode
	if leftNode.Left != nil {
		recursiveLinkPairs(leftNode.Left, leftNode.Right)
		recursiveLinkPairs(rightNode.Left, rightNode.Right)
		recursiveLinkPairs(leftNode.Right, rightNode.Left)
	}
}
