package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	currNode, evenListHead := head, head.Next
	for currNode.Next != nil {
		currNode.Next, currNode = currNode.Next.Next, currNode.Next
	}

	for currNode = head; currNode.Next != nil; currNode = currNode.Next {
	}
	currNode.Next = evenListHead

	return head
}
