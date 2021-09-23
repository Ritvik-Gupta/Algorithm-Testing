package lru_cache

type LRUNode struct {
	key, value int
	next, prev *LRUNode
}

func (node *LRUNode) evict() {
	nextNode, prevNode := node.next, node.prev
	node.next, node.prev = nil, nil
	nextNode.prev, prevNode.next = prevNode, nextNode
}

func (listHead *LRUNode) addNext(node *LRUNode) {
	nextNode := listHead.next
	node.prev, node.next = listHead, nextNode
	nextNode.prev, listHead.next = node, node
}

type LRUCache struct {
	capacity     int
	head         *LRUNode
	tail         *LRUNode
	recordedKeys map[int]*LRUNode
}

func Constructor(capacity int) LRUCache {
	head, tail := &LRUNode{}, &LRUNode{}
	head.next = tail
	tail.prev = head
	return LRUCache{capacity, head, tail, make(map[int]*LRUNode)}
}

func (cache *LRUCache) Get(key int) int {
	if cachedNode, ok := cache.recordedKeys[key]; ok {
		cachedNode.evict()
		cache.head.addNext(cachedNode)
		return cachedNode.value
	}
	return -1
}

func (cache *LRUCache) Put(key, value int) {
	if cachedNode, ok := cache.recordedKeys[key]; ok {
		cachedNode.evict()
		cachedNode.value = value
		cache.head.addNext(cachedNode)
	} else {
		newNode := &LRUNode{key: key, value: value}
		cache.recordedKeys[key] = newNode
		cache.head.addNext(newNode)

		if len(cache.recordedKeys) > cache.capacity {
			tailPrevNode := cache.tail.prev
			tailPrevNode.evict()
			delete(cache.recordedKeys, tailPrevNode.key)
		}
	}
}
