package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	turtle, rabbit := head, head
	var reversed *ListNode

	for rabbit.Next != nil && rabbit.Next.Next != nil {
		rabbit = rabbit.Next.Next
		turtle.Next, reversed, turtle = reversed, turtle, turtle.Next
	}
	turtle.Next, reversed, turtle = reversed, turtle, turtle.Next

	maxSum := 0
	for ; turtle != nil && reversed != nil; turtle, reversed = turtle.Next, reversed.Next {
		fmt.Println(turtle.Val, reversed.Val)
		sum := turtle.Val + reversed.Val
		if maxSum < sum {
			maxSum = sum
		}
	}
	return maxSum
}
