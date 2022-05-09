package medium

type NestedInteger struct {
}

func (nestedInteger NestedInteger) IsInteger() bool {
	panic("")
}
func (nestedInteger NestedInteger) GetInteger() int {
	panic("")
}
func (nestedInteger NestedInteger) GetList() []*NestedInteger {
	panic("")
}

type NestedIterator struct {
	generatorChannel  chan int
	hasNextGeneration bool
}

func flattenNestingGenerator(nestedList []*NestedInteger, generatorChannel chan int, hasNextGen *bool) {
	for _, nestedInteger := range nestedList {
		if nestedInteger.IsInteger() {
			generatorChannel <- nestedInteger.GetInteger()
		} else {
			flattenNestingGenerator(nestedInteger.GetList(), generatorChannel, new(bool))
		}
	}
	*hasNextGen = false
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	obj := NestedIterator{generatorChannel: make(chan int), hasNextGeneration: true}
	go flattenNestingGenerator(nestedList, obj.generatorChannel, &obj.hasNextGeneration)
	return &obj
}

func (nestedIterator *NestedIterator) Next() int {
	return <-nestedIterator.generatorChannel
}

func (nestedIterator *NestedIterator) HasNext() bool {
	return nestedIterator.hasNextGeneration
}
