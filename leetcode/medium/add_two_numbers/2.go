package medium

func defaultListNode(next *ListNode) *ListNode {
	return &ListNode{Val: 0, Next: next}
}

func addLists(listA *ListNode, listB *ListNode) *ListNode {
	if listA == nil && listB == nil {
		return nil
	}

	nextNode := addLists(listA.Next, listB.Next)
	value := listA.Val + listB.Val
	if nextNode != nil && nextNode.Val >= 10 {
		nextNode.Val -= 10
		value += 1
	}
	return &ListNode{Val: value, Next: nextNode}
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	numNodesList1 := 0
	for ptr := l1; ptr != nil; ptr, numNodesList1 = ptr.Next, numNodesList1+1 {
	}

	numNodesList2 := 0
	for ptr := l2; ptr != nil; ptr, numNodesList2 = ptr.Next, numNodesList2+1 {
	}

	listA, listB := defaultListNode(l1), defaultListNode(l2)
	if numNodesList1 < numNodesList2 {
		listA, listB = listB, listA
		numNodesList1, numNodesList2 = numNodesList2, numNodesList1
	}
	for i := numNodesList2; i < numNodesList1; i++ {
		listB = defaultListNode(listB)
	}

	resultList := addLists(listA, listB)
	if resultList.Val == 0 {
		resultList = resultList.Next
	}
	return resultList
}
