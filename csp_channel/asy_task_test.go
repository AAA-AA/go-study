package cspchannel_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
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

//串行执行
func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func TestAsynService(t *testing.T) {
	retCh := AsycService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 1)
}
