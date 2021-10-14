package clone_graph

type Node struct {
	Val       int
	Neighbors []*Node
}

const MAX_NUM_NODES_IN_GRAPH = 100

type NodePair struct {
	originalNode *Node
	clonedNode   *Node
}

func newNodePair(original *Node, cloned *Node)NodePair {
	cloned.Val = original.Val
	cloned.Neighbors = make([]*Node, len(original.Neighbors))

	return NodePair{original, cloned}
}

func CloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	nodeClone := &Node{}

	nodePairQueue := make(chan NodePair, MAX_NUM_NODES_IN_GRAPH)
	nodePairQueue <- newNodePair(node,nodeClone)
	visitedNodes := make(map[int]*Node)
	visitedNodes[node.Val] = nodeClone

iterationOverGraph:
	for {
		select {
		case nodePair := <-nodePairQueue:
			for idx, neighbour := range nodePair.originalNode.Neighbors {
				var neighbourClone *Node
				if cachedNode, ok := visitedNodes[neighbour.Val]; ok {
					neighbourClone = cachedNode
				} else {
					neighbourClone = &Node{}
					nodePairQueue <- newNodePair(neighbour, neighbourClone)
					visitedNodes[neighbourClone.Val] = neighbourClone
				}
				nodePair.clonedNode.Neighbors[idx] = neighbourClone
			}
		default:
			break iterationOverGraph
		}
	}
	close(nodePairQueue)

	return nodeClone
}
