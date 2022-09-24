package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recurTreePath(root *TreeNode, targetSumLeft int, path []int, result *[][]int) {
	targetSumLeft -= root.Val
	path = append(path, root.Val)

	if root.Left == nil && root.Right == nil && targetSumLeft == 0 {
		foundPath := make([]int, 0, len(path))
		for _, val := range path {
			foundPath = append(foundPath, val)
		}
		*result = append(*result, foundPath)
	}

	if root.Left != nil {
		recurTreePath(root.Left, targetSumLeft, path, result)
	}
	if root.Right != nil {
		recurTreePath(root.Right, targetSumLeft, path, result)
	}

	path = path[:len(path)-1]
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	recurTreePath(root, targetSum, []int{}, &result)
	return result
}
