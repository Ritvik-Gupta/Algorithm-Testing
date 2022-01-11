package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursiveVisitNodes(node *TreeNode, currentNum int, finalSum *int) {
	if node == nil {
		return
	}

	currentNum = (currentNum << 1) + node.Val
	if node.Left == nil && node.Right == nil {
		*finalSum += currentNum
	}
	recursiveVisitNodes(node.Left, currentNum, finalSum)
	recursiveVisitNodes(node.Right, currentNum, finalSum)
}

func sumRootToLeaf(root *TreeNode) int {
	finalSum := 0
	recursiveVisitNodes(root, 0, &finalSum)
	return finalSum
}
