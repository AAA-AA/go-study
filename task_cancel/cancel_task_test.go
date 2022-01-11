package taskcancel_test

import (
	"fmt"
	"testing"
	"time"
)

func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func isCanceled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCanceled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancel")
		}(i, cancelChan)
	}
	cancel_1(cancelChan)
	time.Sleep(time.Second * 1)
}
