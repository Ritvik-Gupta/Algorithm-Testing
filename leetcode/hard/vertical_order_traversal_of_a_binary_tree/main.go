package hard

import (
	"container/heap"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type RHElement struct {
	row   uint
	store int
}

type RowWiseHeap []*RHElement

func (colHeap RowWiseHeap) Len() int { return len(colHeap) }

func (colHeap RowWiseHeap) Less(i, j int) bool {
	if colHeap[i].row == colHeap[j].row {
		return colHeap[i].store < colHeap[j].store
	}
	return colHeap[i].row < colHeap[j].row
}

func (colHeap RowWiseHeap) Swap(i, j int) { colHeap[i], colHeap[j] = colHeap[j], colHeap[i] }

func (colHeap *RowWiseHeap) Push(element interface{}) {
	*colHeap = append(*colHeap, element.(*RHElement))
}

func (colHeap *RowWiseHeap) Pop() interface{} {
	element := (*colHeap)[colHeap.Len()-1]
	*colHeap = (*colHeap)[0 : colHeap.Len()-1]
	return element
}

func (colHeap *RowWiseHeap) collect() []int {
	result := make([]int, 0, colHeap.Len())
	for i := colHeap.Len(); i > 0; i-- {
		result = append(result, heap.Pop(colHeap).(*RHElement).store)
	}
	return result
}

func recordsTraversal(node *TreeNode, row uint, column int, records map[int]*RowWiseHeap) {
	if node == nil {
		return
	}

	recordsTraversal(node.Left, row+1, column-1, records)
	recordsTraversal(node.Right, row+1, column+1, records)

	if _, ok := records[column]; !ok {
		records[column] = &RowWiseHeap{}
		heap.Init(records[column])
	}
	heap.Push(records[column], &RHElement{row: row, store: node.Val})
}

func verticalTraversal(root *TreeNode) [][]int {
	verticalOrderRecords := map[int]*RowWiseHeap{}
	recordsTraversal(root, 0, 0, verticalOrderRecords)

	keys := make([]int, 0, len(verticalOrderRecords))
	for key := range verticalOrderRecords {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	results := make([][]int, 0, len(keys))
	for _, key := range keys {
		results = append(results, verticalOrderRecords[key].collect())
	}
	return results
}
