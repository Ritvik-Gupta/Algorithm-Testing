package medium

var BUFFER_LINK_NODE = Node{Val: -101}

func connect1(root *Node) *Node {
	for levelHead := root; levelHead != nil; {
		connectionPointer := &BUFFER_LINK_NODE.Next
		for currNode := levelHead; currNode != nil; currNode = currNode.Next {
			if currNode.Left != nil {
				*connectionPointer = currNode.Left
				connectionPointer = &currNode.Left.Next
			}
			if currNode.Right != nil {
				*connectionPointer = currNode.Right
				connectionPointer = &currNode.Right.Next
			}
		}
		levelHead, BUFFER_LINK_NODE.Next = BUFFER_LINK_NODE.Next, nil
	}
	return root
}
