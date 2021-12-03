package populating_next_right_pointers_in_each_node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	for levelHead := root; levelHead != nil; levelHead = levelHead.Left {
		for currNode := levelHead; currNode != nil; currNode = currNode.Next {
			if currNode.Left != nil {
				currNode.Left.Next = currNode.Right
			}
			if currNode.Right != nil && currNode.Next != nil {
				currNode.Right.Next = currNode.Next.Left
			}
		}

	}
	return root
}
