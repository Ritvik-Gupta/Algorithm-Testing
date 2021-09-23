package wordBreak

import (
	"sync"
	"time"
)

type wordBreakStore struct {
	str         string
	wordDict    []string
	workerGroup *sync.WaitGroup
}

type signal struct{}

func (wbs *wordBreakStore) segmentationWorker(signalChannel chan<- signal, strPtr int) {
	defer wbs.workerGroup.Done()

	if strPtr == len(wbs.str) {
		defer func() { recover() }()
		signalChannel <- signal{}
		time.Sleep(time.Millisecond)
		return
	}

	for _, word := range wbs.wordDict {
		hasSegment := true
		for i := 0; i < len(word); i++ {
			if strPtr+i == len(wbs.str) || wbs.str[strPtr+i] != word[i] {
				hasSegment = false
				break
			}
		}
		if hasSegment {
			wbs.workerGroup.Add(1)
			go wbs.segmentationWorker(signalChannel, strPtr+len(word))
		}
	}
}

func WordBreak(str string, wordDict []string) (canBeSegmented bool) {
	wbs := &wordBreakStore{
		str:         str,
		wordDict:    wordDict,
		workerGroup: &sync.WaitGroup{},
	}
	signalChannel := make(chan signal)

	go func() {
		for range signalChannel {
			canBeSegmented = true
			close(signalChannel)
		}
	}()

	wbs.workerGroup.Add(1)
	go wbs.segmentationWorker(signalChannel, 0)
	wbs.workerGroup.Wait()
	return
}
