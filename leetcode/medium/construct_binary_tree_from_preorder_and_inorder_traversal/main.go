package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeStack []*TreeNode

func (stack *NodeStack) push(node *TreeNode) {
	*stack = append(*stack, node)
}

func (stack NodeStack) top() *TreeNode {
	return stack[len(stack)-1]
}

func (stack *NodeStack) pop() *TreeNode {
	topNode := stack.top()
	*stack = (*stack)[:len(*stack)-1]
	return topNode
}

const OUT_OF_BOUNDS_VALUE = 3001

func buildTree(preorder []int, inorder []int) *TreeNode {
	BUFFER_NODE := TreeNode{Val: OUT_OF_BOUNDS_VALUE}

	preorderIdx, inorderIdx, nextIsLeftChild := 0, 0, true
	nodeStack, currNode := NodeStack{&BUFFER_NODE}, &BUFFER_NODE

	for preorderIdx < len(preorder) {
		if nodeStack.top().Val == inorder[inorderIdx] {
			currNode = nodeStack.pop()
			nextIsLeftChild = false
			inorderIdx++
			continue
		}

		if nextIsLeftChild {
			currNode.Left = &TreeNode{Val: preorder[preorderIdx]}
			currNode = currNode.Left
		} else {
			currNode.Right = &TreeNode{Val: preorder[preorderIdx]}
			currNode = currNode.Right
			nextIsLeftChild = true
		}
		preorderIdx++
		nodeStack.push(currNode)
	}
	return BUFFER_NODE.Left
}
