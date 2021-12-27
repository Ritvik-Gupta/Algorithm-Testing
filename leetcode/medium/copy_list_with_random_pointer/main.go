package medium

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

const MAX_LIST_SIZE = 1000

type NodePair struct {
	original *Node
	cloned   *Node
}

func newNodePair(original *Node, cloned *Node) NodePair {
	cloned.Val = original.Val
	return NodePair{original, cloned}
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	headClone := &Node{}

	nodePairQueue := make(chan NodePair, MAX_LIST_SIZE)
	nodePairQueue <- newNodePair(head, headClone)
	recordedNodes := make(map[*Node]*Node)
	recordedNodes[head] = headClone

fromCloneOverList:
	for {
		select {
		case nodePair := <-nodePairQueue:
			if nextNode := nodePair.original.Next; nextNode != nil {
				var nextNodeClone *Node
				if cachedNode, ok := recordedNodes[nextNode]; ok {
					nextNodeClone = cachedNode
				} else {
					nextNodeClone = &Node{}
					nodePairQueue <- newNodePair(nextNode, nextNodeClone)
					recordedNodes[nextNode] = nextNodeClone
				}
				nodePair.cloned.Next = nextNodeClone
			}
			if randomNode := nodePair.original.Random; randomNode != nil {
				var randomNodeClone *Node
				if cachedNode, ok := recordedNodes[randomNode]; ok {
					randomNodeClone = cachedNode
				} else {
					randomNodeClone = &Node{}
					nodePairQueue <- newNodePair(randomNode, randomNodeClone)
					recordedNodes[randomNode] = randomNodeClone
				}
				nodePair.cloned.Random = randomNodeClone
			}
		default:
			break fromCloneOverList
		}
	}

	close(nodePairQueue)
	return headClone
}
