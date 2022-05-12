package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Using Double Pointers to deal with Linked List `Links`
and reducing the overhead for memory node swapping and shifting

Given a Linked List staring at head, labelling the links :

	Current Link
		 v
	head -> 1 -> 4 -> 3 -> 2 -> 5 -> 2
L:		 @1	  @2   @3   @4	 @5	  @6

	Partition Link = @2
	After Partition Link is decided first Anomaly is at Link @4 ( N(@4)=2 < x=3 )
	Connect @2 to N(@4)=2, @4 to N(@2)=4, @4 to N(@5)=5

	So resuling Linked List structure is :

	head -> 1 -> 2 -> 4 -> 3 -> 5 -> 2
*/
func partition(head *ListNode, x int) *ListNode {
	currLink := &head
	for (*currLink) != nil {
		if (*currLink).Val >= x {
			break
		}
		currLink = &(*currLink).Next
	}
	partitionLink := currLink

	for *currLink != nil {
		currNode := *currLink
		if currNode.Val < x {
			*currLink = currNode.Next
			currNode.Next = *partitionLink

			*partitionLink = currNode
			partitionLink = &(*partitionLink).Next
		} else {
			currLink = &(*currLink).Next
		}
	}

	return head
}
