package hard

const MAX_POSSIBLE_STORE = 50001

var DEFAULT_EMPTY_KEY = ""

type Node struct {
	key        *string
	store      uint
	next, prev *Node
}

func NewNode(key *string, store uint) *Node {
	return &Node{key: key, store: store}
}

func (node *Node) inc() {
	node.store += 1
	prevNode, _ := node.evict()
	for prevNode.store <= node.store {
		prevNode = prevNode.prev
	}
	prevNode.addNext(node)
}

func (node *Node) dec() {
	node.store -= 1
	_, nextNode := node.evict()
	if node.store == 0 {
		return
	}

	for nextNode.store >= node.store {
		nextNode = nextNode.next
	}
	nextNode.addPrev(node)
}

func (node *Node) evict() (prevNode, nextNode *Node) {
	prevNode, nextNode = node.prev, node.next
	node.prev, node.next = nil, nil
	nextNode.prev, prevNode.next = prevNode, nextNode
	return
}

func (node *Node) addNext(newNode *Node) {
	nextNode := node.next
	newNode.prev, node.next = node, newNode
	nextNode.prev, newNode.next = newNode, nextNode
}

func (node *Node) addPrev(newNode *Node) {
	prevNode := node.prev
	node.prev, newNode.next = newNode, node
	newNode.prev, prevNode.next = prevNode, newNode
}

type AllOne struct {
	head, tail    *Node
	recordedNodes map[string]*Node
}

func Constructor() AllOne {
	head, tail := NewNode(&DEFAULT_EMPTY_KEY, MAX_POSSIBLE_STORE), NewNode(&DEFAULT_EMPTY_KEY, 0)
	tail.prev, head.next = head, tail
	return AllOne{head: head, tail: tail, recordedNodes: make(map[string]*Node)}
}

func (register *AllOne) Inc(key string) {
	// println("Inc ", key)
	if recordedNode, ok := register.recordedNodes[key]; ok {
		recordedNode.inc()
	} else {
		newNode := NewNode(&key, 1)
		register.tail.addPrev(newNode)
		register.recordedNodes[key] = newNode
	}
}

func (register *AllOne) Dec(key string) {
	// println("Dec ", key)
	recordedNode := register.recordedNodes[key]
	recordedNode.dec()
	if recordedNode.store == 0 {
		delete(register.recordedNodes, key)
	}
}

func (register *AllOne) GetMaxKey() string {
	// println("Get Max Key")
	return *register.head.next.key
}

func (register *AllOne) GetMinKey() string {
	// println("Get Min Key")
	return *register.tail.prev.key
}

/*
["AllOne","inc a","inc b","inc b","inc c","inc c","inc c","dec b", "dec b","getMinKey","dec a","getMaxKey","getMinKey"]

[null,null,null,null,null,null,null,null,null,"b",null,"c","b"]

[null,null,null,null,null,null,null,null,null,"a",null,"c","c"]
*/
