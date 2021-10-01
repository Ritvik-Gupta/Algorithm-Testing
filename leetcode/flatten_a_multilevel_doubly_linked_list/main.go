package flatten_a_multilevel_doubly_linked_list

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flattenAndEndNode(root *Node) *Node {
	if root.Next == nil && root.Child == nil {
		return root
	}

	if root.Next != nil && root.Child != nil {
		endOfNext, endOfChild := flattenAndEndNode(root.Next), flattenAndEndNode(root.Child)
		endOfChild.Next, root.Next.Prev = root.Next, endOfChild
		root.Next, root.Child.Prev = root.Child, root
		root.Child = nil
		return endOfNext
	}

	if root.Child != nil {
		root.Next, root.Child.Prev = root.Child, root
		root.Child = nil
	}
	return flattenAndEndNode(root.Next)
}

func flatten(root *Node) *Node {
	if root != nil {
		flattenAndEndNode(root)
	}
	return root
}
