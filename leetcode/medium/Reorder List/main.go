package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prevNode *ListNode
	currNode := head

	for currNode != nil {
		nextNode := currNode.Next
		currNode.Next = prevNode

		prevNode, currNode = currNode, nextNode
	}

	return prevNode
}

func reorderList(head *ListNode) {
	turtle, rabbit := head, head

	for rabbit != nil {
		turtle = turtle.Next
		rabbit = rabbit.Next
		if rabbit == nil {
			continue
		}
		rabbit = rabbit.Next
	}

	fstHalfNode, scnHalfNode := head, reverseList(turtle)

	for fstHalfNode != nil {
		nextNode := fstHalfNode.Next
		fstHalfNode.Next = scnHalfNode

		fstHalfNode, scnHalfNode = scnHalfNode, nextNode
	}
}
