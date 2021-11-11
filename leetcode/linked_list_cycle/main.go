package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	turtle, hare := head, head
	for {
		if turtle == nil {
			return false
		}
		turtle = turtle.Next

		hare = hare.Next
		if hare == nil {
			return false
		}
		hare = hare.Next
		if hare == nil {
			return false
		}

		if turtle == hare {
			return true
		}
	}
}
