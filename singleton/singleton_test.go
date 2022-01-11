package singleton_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singelton struct{}

var instance *Singelton

var once sync.Once

func GetInstance() *Singelton {
	once.Do(func() {
		fmt.Println("Create Obj")
		instance = new(Singelton)
	})
	return instance
}

func TestSingleton(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetInstance()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
