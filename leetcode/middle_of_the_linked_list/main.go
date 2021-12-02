package middle_of_the_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	turtle, rabbit := head, head
	for rabbit != nil && rabbit.Next != nil {
		turtle = turtle.Next
		rabbit = rabbit.Next.Next
	}
	return turtle
}
