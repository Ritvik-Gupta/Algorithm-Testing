package medium

func getAndMoveNext(nodePtr **ListNode) int {
	if node := *nodePtr; node == nil {
		return 0
	} else {
		*nodePtr = node.Next
		return node.Val
	}
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p1, p2, p3, carry := l1, l2, result, 0

	for p1 != nil || p2 != nil || carry != 0 {
		val1, val2 := getAndMoveNext(&p1), getAndMoveNext(&p2)
		sum := val1 + val2 + carry
		if sum >= 10 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}
		p3.Next = &ListNode{Val: sum}
		p3 = p3.Next
	}

	return result.Next
}
