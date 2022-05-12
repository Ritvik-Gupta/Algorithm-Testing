package medium

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func allSectionsLeafs(store *Node) bool {
	return store.TopLeft.IsLeaf && store.TopRight.IsLeaf &&
		store.BottomLeft.IsLeaf && store.BottomRight.IsLeaf
}

func allSectionsSameValue(store *Node) bool {
	randomSectionValue := store.TopLeft.Val
	return store.TopRight.Val == randomSectionValue &&
		store.BottomLeft.Val == randomSectionValue &&
		store.BottomRight.Val == randomSectionValue
}

func pruneSections(store *Node) {
	store.TopLeft = nil
	store.TopRight = nil
	store.BottomLeft = nil
	store.BottomRight = nil
}

func intersect(store, traverse *Node) *Node {
	if store.IsLeaf && traverse.IsLeaf {
		store.Val = store.Val || traverse.Val
		return store
	}

	if traverse.IsLeaf && !store.IsLeaf {
		return intersect(traverse, store)
	}

	if store.IsLeaf && !traverse.IsLeaf {
		if store.Val == true {
			return store
		} else {
			return traverse
		}
	}

	if !store.IsLeaf && !traverse.IsLeaf {
		store.TopLeft = intersect(store.TopLeft, traverse.TopLeft)
		store.TopRight = intersect(store.TopRight, traverse.TopRight)
		store.BottomLeft = intersect(store.BottomLeft, traverse.BottomLeft)
		store.BottomRight = intersect(store.BottomRight, traverse.BottomRight)

		if allSectionsSameValue(store) && allSectionsLeafs(store) {
			store.IsLeaf = true
			store.Val = store.TopLeft.Val
			pruneSections(store)
		}

		return store
	}

	return store
}
