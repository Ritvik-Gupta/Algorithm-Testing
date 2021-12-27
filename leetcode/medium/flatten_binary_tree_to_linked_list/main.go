package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flattenAndRightmostChild(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		return root
	} else if root.Left != nil && root.Right != nil {
		rightmostOfLeft := flattenAndRightmostChild(root.Left)
		rightmostOfRight := flattenAndRightmostChild(root.Right)

		rightmostOfLeft.Right = root.Right
		root.Right, root.Left = root.Left, nil
		return rightmostOfRight
	} else {
		if root.Right == nil {
			root.Right, root.Left = root.Left, nil
		}
		return flattenAndRightmostChild(root.Right)
	}
}

func flatten2(root *TreeNode) {
	if root != nil {
		flattenAndRightmostChild(root)
	}
}
