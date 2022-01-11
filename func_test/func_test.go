package func_test

import (
	"fmt"
	"math/rand"
	"testing"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
}

func Sum(ops ...int) int {
	ret := 0
	for _, v := range ops {
		ret = ret + v
	}
	return ret
}

func TestSum(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resource.")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	panic("err")
}
