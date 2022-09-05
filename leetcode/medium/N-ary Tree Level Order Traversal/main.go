package medium

type Node struct {
	Val      int
	Children []*Node
}

func levelOrderRecur(root *Node, depth int, result *[][]int) {
	if depth == len(*result) {
		*result = append(*result, []int{})
	}

	(*result)[depth] = append((*result)[depth], root.Val)
	for _, childNode := range root.Children {
		levelOrderRecur(childNode, depth+1, result)
	}
}

func levelOrder(root *Node) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}
	levelOrderRecur(root, 0, &result)
	return result
}
