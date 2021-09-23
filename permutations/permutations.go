package permutations

import (
	"sync"
)

/*
? Given an array nums of distinct integers, return all the possible permutations.
? You can return the answer in any order.
*/

type _Permutations struct {
	result       [][]int
	workerWriter *sync.Mutex
	workerGroup  *sync.WaitGroup
}

func (permutations *_Permutations) permuteWorker(prevNums, leftNums []int) {
	if len(leftNums) == 0 {
		permutations.workerWriter.Lock()
		permutations.result = append(permutations.result, prevNums)
		permutations.workerWriter.Unlock()
	}

	for i, stored := range leftNums {
		newLeftNums := make([]int, i)
		copy(newLeftNums, leftNums[:i])

		permutations.workerGroup.Add(1)
		go permutations.permuteWorker(
			append(prevNums, stored),
			append(newLeftNums, leftNums[i+1:]...),
		)
	}

	permutations.workerGroup.Done()
}

func Permute(nums []int) [][]int {
	permutations := _Permutations{
		result:       [][]int{},
		workerWriter: &sync.Mutex{},
		workerGroup:  &sync.WaitGroup{},
	}

	permutations.workerGroup.Add(1)
	go permutations.permuteWorker([]int{}, nums)
	permutations.workerGroup.Wait()

	return permutations.result
}
