package decoding

import (
	"sync"
	"time"
)

type signal struct{}

type _CharSequence struct{ first, second byte }

func (seq *_CharSequence) isValid() bool {
	if fd, sd := seq.first, seq.second; '0' <= sd && sd <= '9' {
		return fd == '1' || (fd == '2' && sd <= '6')
	} else {
		return '1' <= fd && fd <= '9'
	}
}

func decodingWorker(
	strPtr *string,
	startPos int,
	decodingPoolGroup *sync.WaitGroup,
	signalChannel chan<- signal,
) {
	defer decodingPoolGroup.Done()
	str := *strPtr

	for i := startPos; i < len(str); i++ {
		if j := i + 1; j < len(str) {
			if isValidSeq := (&_CharSequence{str[i], str[j]}).isValid(); isValidSeq {
				decodingPoolGroup.Add(1)
				go decodingWorker(strPtr, i+2, decodingPoolGroup, signalChannel)
			}
		}

		if isValidSeq := (&_CharSequence{first: str[i]}).isValid(); !isValidSeq {
			return
		}
	}
	signalChannel <- signal{}
}

func NumDecodings(str string) (numDecodingWays int) {
	decodingPoolGroup := &sync.WaitGroup{}
	signalChannel := make(chan signal)

	go func() {
		for range signalChannel {
			numDecodingWays++
		}
	}()

	decodingPoolGroup.Add(1)
	go decodingWorker(&str, 0, decodingPoolGroup, signalChannel)
	decodingPoolGroup.Wait()
	time.Sleep(time.Millisecond)

	return
}
