package recover_a_tree_from_preorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeBuilder struct {
	root      *TreeNode
	currNode  *TreeNode
	heirarchy map[*TreeNode]*TreeNode
}

func newTreeBuilder() *TreeBuilder {
	root := &TreeNode{Left: &TreeNode{}}
	treeBuilder := TreeBuilder{root, root, make(map[*TreeNode]*TreeNode)}
	treeBuilder.createChild(0, LEFT)
	return &treeBuilder
}

type TreeChild uint

const (
	LEFT TreeChild = iota
	RIGHT
)

func (builder *TreeBuilder) goToParent() {
	builder.currNode = builder.heirarchy[builder.currNode]
}

func (builder *TreeBuilder) createChild(value int, toCreate TreeChild) {
	var childNode **TreeNode
	switch toCreate {
	case LEFT:
		childNode = &builder.currNode.Left
	case RIGHT:
		childNode = &builder.currNode.Right
	}
	*childNode = &TreeNode{Val: value}
	builder.heirarchy[*childNode], builder.currNode = builder.currNode, *childNode
}

const DIGIT_OFFSET = int('0')
const SEPARATION = '-'

func recoverFromPreorder(traversal string) *TreeNode {
	treeBuilder := newTreeBuilder()

	for currLevel, i := 0, 0; i < len(traversal); {
		value, numSeparations := 0, 0
		for ; i < len(traversal) && traversal[i] == SEPARATION; i++ {
			numSeparations++
		}
		for ; i < len(traversal) && traversal[i] != SEPARATION; i++ {
			value = value*10 + int(traversal[i]) - DIGIT_OFFSET
		}

		toCreate := LEFT
		for j := numSeparations; j <= currLevel; j++ {
			toCreate = RIGHT
			treeBuilder.goToParent()
		}

		treeBuilder.createChild(value, toCreate)
		currLevel = numSeparations
	}

	return treeBuilder.root.Right
}
