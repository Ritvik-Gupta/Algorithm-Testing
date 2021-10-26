package the_skyline_problem

type Vector struct {
	x, y int
}

type Rectangle struct {
	start, end, height int
}

type VectorChain struct {
	*Vector
	next, prev *VectorChain
}

type Polygon struct {
	head, tail *VectorChain
}

func newPolygon() Polygon {
	head, tail := &VectorChain{}, &VectorChain{}
	tail.prev, head.next = head, tail
	return Polygon{head, tail}
}

func (polygon *Polygon) pushBack(vector *Vector) {
	newVec, tailPrevVec := &VectorChain{Vector: vector}, polygon.tail.prev
	newVec.prev, newVec.next = tailPrevVec, polygon.tail
	tailPrevVec.next, polygon.tail.prev = newVec, newVec
}

func (polygon *Polygon) add(rect *Rectangle) {
	if polygon.head == polygon.tail {
		polygon.pushBack(&Vector{rect.start, rect.height})
		polygon.pushBack(&Vector{rect.end, 0})
		return
	}

	currVec := polygon.head
	for currVec != polygon.tail {

		currVec = currVec.next
	}
}

func getSkyline(buildings [][]int) [][]int {
	panic("")
}
