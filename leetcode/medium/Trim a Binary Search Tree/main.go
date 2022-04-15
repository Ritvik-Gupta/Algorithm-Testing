package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low, high int) *TreeNode {
	var trimRecur func(**TreeNode)

	trimRecur = func(root **TreeNode) {
	CHECK_ROOT_VALIDITY:

		if *root == nil {
			return
		}

		if (*root).Val < low {
			*root = (*root).Right
			goto CHECK_ROOT_VALIDITY
		} else if (*root).Val > high {
			*root = (*root).Left
			goto CHECK_ROOT_VALIDITY
		}

		trimRecur(&(*root).Left)
		trimRecur(&(*root).Right)
	}

	trimRecur(&root)
	return root
}
