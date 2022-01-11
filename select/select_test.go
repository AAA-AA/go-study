package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done")
}

func AsycService() chan string {
	retCh := make(chan string, 1) //buffer channel
	//retCh := make(chan string) 阻塞类型chanel
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	select {
	case res := <-AsycService():
		t.Log(res)
	case <-time.After(time.Millisecond * 100):
		t.Error("timeout")
	}

}
