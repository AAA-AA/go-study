package objpool_test

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type ReusableObj struct{}

type ObjPool struct {
	bufChan chan *ReusableObj
}

func NewObjPool(numOfObj int) *ObjPool {
	pool := ObjPool{}
	pool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		pool.bufChan <- &ReusableObj{}
	}

	return &pool

}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

func (p *ObjPool) Release(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	for i := 0; i < 11; i++ {
		if v, error := pool.GetObj(time.Second * 1); error != nil {
			t.Error(error)
		} else {
			fmt.Printf("%T\n", v)
			if err := pool.Release(v); err != nil {
				t.Error(err)
			}
		}
	}
}
