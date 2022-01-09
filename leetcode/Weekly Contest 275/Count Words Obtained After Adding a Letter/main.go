package main

import (
	"sort"
	"strings"
)

type TrieNode struct {
	store    string
	children map[byte]*TrieNode
}

func newTrieNode() *TrieNode {
	return &TrieNode{children: make(map[byte]*TrieNode)}
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func wordCount(startWords []string, targetWords []string) int {
	trie := newTrieNode()
	maxTrieLen := 0

	for _, word := range targetWords {
		word = sortString(word)

		currNode := trie
		for i := 0; i < len(word); i++ {
			if _, ok := currNode.children[word[i]]; !ok {
				currNode.children[word[i]] = newTrieNode()
			}
			currNode = currNode.children[word[i]]
		}
		currNode.store = word

		if maxTrieLen < len(word) {
			maxTrieLen = len(word)
		}
	}

	allTargetWordsFound := make(map[string]struct{})
	for _, word := range startWords {
		word = sortString(word)

		for i := 0; i <= len(word); i++ {
			targetWordsFound := checkCombination(word, i, trie, maxTrieLen)
			for targetWordFound := range targetWordsFound {
				allTargetWordsFound[targetWordFound] = struct{}{}
			}
		}
	}

	return len(allTargetWordsFound)
}

func checkCombination(word string, spaceAt int, trie *TrieNode, maxTrieLen int) map[string]struct{} {
	targetWordsFound := make(map[string]struct{})

	if len(word)+1 > maxTrieLen {
		return targetWordsFound
	}

	currNode := trie
	for i := 0; i < spaceAt; i++ {
		if childNode, ok := currNode.children[word[i]]; !ok {
			return targetWordsFound
		} else {
			currNode = childNode
		}
	}

	possibleNodes := make([]*TrieNode, 0, len(currNode.children))
	for _, childNode := range currNode.children {
		possibleNodes = append(possibleNodes, childNode)
	}

	for i := spaceAt; i < len(word); i++ {
		for j := 0; j < len(possibleNodes); j++ {
			if possibleNodes[j] == nil {
				continue
			}
			if childNode, ok := possibleNodes[j].children[word[i]]; ok {
				possibleNodes[j] = childNode
			} else {
				possibleNodes[j] = nil
			}
		}
	}

	for j := 0; j < len(possibleNodes); j++ {
		if possibleNodes[j] == nil {
			continue
		}
		if possibleNodes[j].store != "" {
			targetWordsFound[possibleNodes[j].store] = struct{}{}
		}
	}
	return targetWordsFound
}
