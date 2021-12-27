package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	tilt := 0
	computeTiltAndFindSum(root, &tilt)
	return tilt
}

func computeTiltAndFindSum(root *TreeNode, totalTilt *int) int {
	if root == nil {
		return 0
	}

	leftSum := computeTiltAndFindSum(root.Left, totalTilt)
	rightSum := computeTiltAndFindSum(root.Right, totalTilt)
	*totalTilt += abs(leftSum - rightSum)

	return leftSum + rightSum + root.Val
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
