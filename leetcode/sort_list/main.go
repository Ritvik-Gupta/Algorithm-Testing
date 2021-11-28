package sort_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	middleNode := fetchMiddleNode(head)
	nextToMiddleNode := middleNode.Next
	middleNode.Next = nil

	return mergeSortedLists(sortList(head), sortList(nextToMiddleNode))
}

func fetchMiddleNode(head *ListNode) (turtle *ListNode) {
	if head == nil {
		return
	}

	turtle, rabbit := head, head
	for rabbit.Next != nil && rabbit.Next.Next != nil {
		turtle = turtle.Next
		rabbit = rabbit.Next.Next
	}
	return
}

func mergeSortedLists(left, right *ListNode) (result *ListNode) {
	if left == nil {
		return right
	} else if right == nil {
		return left
	}

	if left.Val > right.Val {
		result = right
		result.Next = mergeSortedLists(left, right.Next)
	} else {
		result = left
		result.Next = mergeSortedLists(left.Next, right)
	}
	return
}
