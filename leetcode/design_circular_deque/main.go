package design_circular_deque

type LinkedNode struct {
	store      int
	next, prev *LinkedNode
}

func NewNode(store int) LinkedNode {
	return LinkedNode{store: store}
}

type MyCircularDeque struct {
	maxSize, totalNodes uint
	head                *LinkedNode
}

func Constructor(maxSize int) MyCircularDeque {
	return MyCircularDeque{maxSize: uint(maxSize)}
}

func (deque *MyCircularDeque) shiftHeadToPrev() {
	if !deque.IsEmpty() {
		deque.head = deque.head.prev
	}
}

func (deque *MyCircularDeque) InsertFront(value int) bool {
	if deque.InsertLast(value) {
		deque.shiftHeadToPrev()
		return true
	}
	return false
}

func (deque *MyCircularDeque) InsertLast(value int) bool {
	if deque.IsFull() {
		return false
	}
	defer func() { deque.totalNodes++ }()

	newNode := NewNode(value)
	if deque.IsEmpty() {
		deque.head = &newNode
		deque.head.next, deque.head.prev = deque.head, deque.head
		return true
	}

	headNode, prevNode := deque.head, deque.head.prev
	newNode.prev, prevNode.next = prevNode, &newNode
	newNode.next, headNode.prev = headNode, &newNode
	return true
}

func (deque *MyCircularDeque) DeleteFront() bool {
	if deque.IsEmpty() {
		return false
	}
	defer func() { deque.totalNodes-- }()

	prevNode, nextNode := deque.head.prev, deque.head.next
	nextNode.prev, prevNode.next = prevNode, nextNode
	deque.head.prev, deque.head.next = nil, nil

	if deque.totalNodes == 1 {
		deque.head = nil
	} else {
		deque.head = nextNode
	}
	return true
}

func (deque *MyCircularDeque) DeleteLast() bool {
	deque.shiftHeadToPrev()
	return deque.DeleteFront()
}

func (deque *MyCircularDeque) GetFront() int {
	if deque.head == nil {
		return -1
	}
	return deque.head.store
}

func (deque *MyCircularDeque) GetRear() int {
	if deque.head == nil {
		return -1
	}
	return deque.head.prev.store
}

func (deque *MyCircularDeque) IsEmpty() bool {
	return deque.totalNodes == 0
}

func (deque *MyCircularDeque) IsFull() bool {
	return deque.totalNodes == deque.maxSize
}

func (deque *MyCircularDeque) Traverse() {
	if deque.head == nil {
		println("List is Empty")
		return
	}

	node := deque.head
	for {
		print(node.store, " -> ")
		node = node.next
		if node == deque.head {
			break
		}
	}
	println()
}
