package balanced_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return findRecursiveHeight(root) != -1
}

func findRecursiveHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := findRecursiveHeight(root.Left), findRecursiveHeight(root.Right)
	if left >= 0 && right >= 0 && isAlmostUnit(left-right) {
		return max(left, right) + 1
	}
	return -1
}

func isAlmostUnit(val int) bool {
	return val == 0 || val == 1 || val == -1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
