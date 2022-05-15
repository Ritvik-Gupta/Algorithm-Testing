package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	treeDepth, leavesSum := 0, 0

	var dfsDeepestLeavesSum func(*TreeNode, int)
	dfsDeepestLeavesSum = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		dfsDeepestLeavesSum(root.Left, depth+1)
		dfsDeepestLeavesSum(root.Right, depth+1)

		if depth > treeDepth {
			treeDepth = depth
			leavesSum = root.Val
		} else if depth == treeDepth {
			leavesSum += root.Val
		}
	}

	dfsDeepestLeavesSum(root, 0)
	return leavesSum
}
