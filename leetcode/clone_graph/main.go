package clone_graph

type Node struct {
	Val       int
	Neighbors []*Node
}

const (
	MAX_NUM_NODES_IN_GRAPH = 100
	GIVEN_MAX_VAL_OF_NODE  = 1000
)

type NodePair struct {
	originalNode *Node
	clonedNode   *Node
}

func newNodePair(original *Node, cloned *Node) NodePair {
	cloned.Val = original.Val
	cloned.Neighbors = make([]*Node, len(original.Neighbors))

	return NodePair{original, cloned}
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	nodeClone := &Node{}

	nodePairQueue := make(chan NodePair, MAX_NUM_NODES_IN_GRAPH)
	nodePairQueue <- newNodePair(node, nodeClone)

	node.Neighbors = append(node.Neighbors, nodeClone)
	node.Val += GIVEN_MAX_VAL_OF_NODE

fromCloneOverGraph:
	for {
		select {
		case nodePair := <-nodePairQueue:
			for idx := 0; idx < len(nodePair.originalNode.Neighbors)-1; idx++ {
				neighbourNode := nodePair.originalNode.Neighbors[idx]
				var neighbourNodeClone *Node
				if neighbourNode.Val > GIVEN_MAX_VAL_OF_NODE {
					neighbourNodeClone = neighbourNode.Neighbors[len(neighbourNode.Neighbors)-1]
				} else {
					neighbourNodeClone = &Node{}
					nodePairQueue <- newNodePair(neighbourNode, neighbourNodeClone)

					neighbourNode.Neighbors = append(neighbourNode.Neighbors, neighbourNodeClone)
					neighbourNode.Val += GIVEN_MAX_VAL_OF_NODE
				}
				nodePair.clonedNode.Neighbors[idx] = neighbourNodeClone
			}
		default:
			break fromCloneOverGraph
		}
	}

	nodeQueue := make(chan *Node, MAX_NUM_NODES_IN_GRAPH)
	nodeQueue <- node

	node.Val -= GIVEN_MAX_VAL_OF_NODE
	node.Neighbors = node.Neighbors[:len(node.Neighbors)-1]

fromCleanUpGraph:
	for {
		select {
		case node := <-nodeQueue:
			for _, neighbourNode := range node.Neighbors {
				if neighbourNode.Val > GIVEN_MAX_VAL_OF_NODE {
					nodeQueue <- neighbourNode

					neighbourNode.Val -= GIVEN_MAX_VAL_OF_NODE
					neighbourNode.Neighbors = neighbourNode.Neighbors[:len(neighbourNode.Neighbors)-1]
				}
			}
		default:
			break fromCleanUpGraph
		}
	}

	close(nodePairQueue)
	return nodeClone
}
