package testpackage_test

import "testing"

func TestImplict(t *testing.T) {
	var a int = 1
	var b int64
	b = int64(a)
	t.Log(a, b)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
}
