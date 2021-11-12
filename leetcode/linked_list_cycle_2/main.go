package linked_list_cycle_2

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) (snail *ListNode) {
	turtle, rabbit := head, head
	for rabbit != nil && rabbit.Next != nil {
		turtle = turtle.Next
		rabbit = rabbit.Next.Next

		if turtle == rabbit {
			snail = head
			for turtle != snail {
				turtle = turtle.Next
				snail = snail.Next
			}
			return
		}
	}
	return
}
