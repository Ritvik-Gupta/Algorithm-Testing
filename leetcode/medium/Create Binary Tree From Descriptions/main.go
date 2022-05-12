package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) (rootNode *TreeNode) {
	treeRecord := make(map[int]*TreeNode)
	parentRecord := make(map[*TreeNode]*TreeNode)

	for idx, description := range descriptions {
		parentValue, childValue, isLeft := description[0], description[1], description[2] == 1
		if _, ok := treeRecord[parentValue]; !ok {
			treeRecord[parentValue] = &TreeNode{Val: parentValue}
		}
		if _, ok := treeRecord[childValue]; !ok {
			treeRecord[childValue] = &TreeNode{Val: childValue}
		}
		parent, child := treeRecord[parentValue], treeRecord[childValue]
		parentRecord[child] = parent

		if isLeft {
			parent.Left = child
		} else {
			parent.Right = child
		}

		if idx == len(descriptions)-1 {
			rootNode = child
			for parentRecord[rootNode] != nil {
				rootNode = parentRecord[rootNode]
			}
		}
	}

	return
}
