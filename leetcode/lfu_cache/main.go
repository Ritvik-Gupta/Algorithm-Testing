package lfu_cache

type LFUStore struct {
	key, value, rank int
}

func createStore(key, value int) *LFUStore {
	return &LFUStore{key: key, value: value, rank: 1}
}

func (store *LFUStore) rankUp() {
	store.rank += 1
}

func (store *LFUStore) updateValue(value int) {
	store.value = value
	store.rankUp()
}

type LFUNode struct {
	store      *LFUStore
	next, prev *LFUNode
}

func (node *LFUNode) isRankToken() bool { return node.store == nil }

func (node *LFUNode) addNext(newNode *LFUNode) {
	nextNode := node.next
	newNode.prev, newNode.next = node, nextNode
	nextNode.prev, node.next = newNode, newNode
}

func (node *LFUNode) tryToDrop() (droppedKey int, hasBeenDropped bool) {
	if node.isRankToken() {
		return
	}
	prevNode, nextNode := node.prev, node.next
	node.prev, node.next = nil, nil
	prevNode.next, nextNode.prev = nextNode, prevNode
	return node.store.key, true
}

type LFUCache struct {
	capacity     int
	rankTokens   []*LFUNode
	recordedKeys map[int]*LFUNode
}

func Constructor(capacity int) LFUCache {
	head, tail := &LFUNode{}, &LFUNode{}
	tail.prev, head.next = head, tail
	return LFUCache{
		capacity:     capacity,
		rankTokens:   []*LFUNode{tail, head},
		recordedKeys: make(map[int]*LFUNode),
	}
}

func (cache *LFUCache) getRankToken(idx int) *LFUNode {
	if tokenLen := len(cache.rankTokens); tokenLen == idx {
		newHead, oldHead := &LFUNode{}, cache.rankTokens[tokenLen-1]
		oldHead.prev, newHead.next = newHead, oldHead
		cache.rankTokens = append(cache.rankTokens, newHead)
	}
	return cache.rankTokens[idx]
}

func (cache *LFUCache) Get(key int) int {
	if cachedNode, ok := cache.recordedKeys[key]; ok {
		cachedNode.tryToDrop()
		cachedNode.store.rankUp()
		cache.getRankToken(cachedNode.store.rank).addNext(cachedNode)
		return cachedNode.store.value
	}
	return -1
}

func (cache *LFUCache) Put(key int, value int) {
	if cache.capacity == 0 {
		return
	}

	if cachedNode, ok := cache.recordedKeys[key]; ok {
		cachedNode.tryToDrop()
		cachedNode.store.updateValue(value)
		cache.getRankToken(cachedNode.store.rank).addNext(cachedNode)
	} else {
		if len(cache.recordedKeys) == cache.capacity {
			for _, rankToken := range cache.rankTokens {
				if droppedKey, hasBeenDropped := rankToken.prev.tryToDrop(); hasBeenDropped {
					delete(cache.recordedKeys, droppedKey)
					break
				}
			}
		}
		newNode := &LFUNode{store: createStore(key, value)}
		cache.recordedKeys[key] = newNode
		cache.getRankToken(newNode.store.rank).addNext(newNode)
	}
}
