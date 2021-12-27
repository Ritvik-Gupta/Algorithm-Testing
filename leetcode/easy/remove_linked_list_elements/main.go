package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

const OUT_OF_BOUNDS_VAL = 51

func removeElements(head *ListNode, val int) *ListNode {
	head = &ListNode{Val: OUT_OF_BOUNDS_VAL, Next: head}
	for currNode := head; currNode.Next != nil; {
		if currNode.Next.Val == val {
			currNode.Next.Next, currNode.Next = nil, currNode.Next.Next
		} else {
			currNode = currNode.Next
		}
	}
	return head.Next
}
