package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	turtle, rabbit := &head, head

	for rabbit != nil && rabbit.Next != nil {
		turtle = &(*turtle).Next
		rabbit = rabbit.Next
		if rabbit != nil {
			rabbit = rabbit.Next
		}
	}

	*turtle = (*turtle).Next

	return head
}
