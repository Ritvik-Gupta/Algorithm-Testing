package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type MonotonicStack []*TreeNode

func newStack(size int) MonotonicStack {
	return make([]*TreeNode, size)
}

func (stack *MonotonicStack) push(value *TreeNode) {
	*stack = append(*stack, value)
}

func (stack *MonotonicStack) pop() *TreeNode {
	stackLength := len(*stack)
	lastValue := (*stack)[stackLength-1]
	*stack = (*stack)[:stackLength-1]
	return lastValue
}

func (stack *MonotonicStack) peek() *TreeNode {
	return (*stack)[len(*stack)-1]
}

func (orderedStack *MonotonicStack) tryMonotonicPush(node *TreeNode) {
	for prevNode := orderedStack.peek(); node.Val > prevNode.Val; {
		lastNode := orderedStack.pop()

		if prevNode = orderedStack.peek(); node.Val > prevNode.Val {
			prevNode.Right = lastNode
		} else {
			node.Left = lastNode
		}
	}
	orderedStack.push(node)
}

var INFINITE_VALUE_NODE = TreeNode{Val: 1001}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	orderedStack := newStack(len(nums))
	orderedStack.push(&INFINITE_VALUE_NODE)

	for _, num := range nums {
		orderedStack.tryMonotonicPush(&TreeNode{Val: num})
	}
	orderedStack.tryMonotonicPush(&INFINITE_VALUE_NODE)
	return INFINITE_VALUE_NODE.Left
}
