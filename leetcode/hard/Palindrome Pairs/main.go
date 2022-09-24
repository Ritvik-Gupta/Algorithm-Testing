package hard

func isPalindrome(word string) bool {
	lo, hi := 0, len(word)-1
	for lo <= hi {
		if word[lo] != word[hi] {
			return false
		}
		lo++
		hi--
	}
	return true
}

type TrieNode struct {
	token      byte
	isEnd      bool
	depth      int
	parentNode *TrieNode
	children   map[byte]*TrieNode
}

func NewNode(token byte, parentNode *TrieNode) *TrieNode {
	depth := 0
	if parentNode != nil {
		depth = parentNode.depth + 1
	}
	return &TrieNode{
		token:      token,
		isEnd:      false,
		depth:      depth,
		parentNode: parentNode,
		children:   make(map[byte]*TrieNode),
	}
}

func (root *TrieNode) insert(word string) *TrieNode {
	currNode := root
	for i := 0; i < len(word); i++ {
		token := word[i]
		if currNode.children[token] == nil {
			currNode.children[token] = NewNode(token, currNode)
		}
		currNode = currNode.children[token]
	}
	currNode.isEnd = true
	return currNode
}

func palindromePairs(words []string) [][]int {
	root := NewNode(' ', nil)
	wordLastNodes := make([]*TrieNode, 0, len(words))

	for _, word := range words {
		lastNode := root.insert(word)
		wordLastNodes = append(wordLastNodes, lastNode)
	}

	for _, lastNode := range wordLastNodes {
		backNode := lastNode
		currNode := root

		for backNode != root {
			currNode = currNode.children[backNode.token]
			if currNode == nil {
				break
			}
			backNode = backNode.parentNode
		}

		if currNode == nil {

		}
	}

	return nil
}
