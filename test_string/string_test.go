package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	//s[1] = '3'  //string是不可变byte slice

	s = "\xE4\xB8\xA5" //string可以存储任何二进制数据
	t.Log(s)
	t.Log(len(s))

	s = "中"
	t.Log(len(s))

	c := []rune(s)
	t.Log(len(c))

	t.Logf("中 unicode %x", c[0])

	t.Logf("中 UTF8 %x", s)

}

func TestStirngFn(t *testing.T) {
	s := "a,b,c"
	parts := strings.Split(s, ",")
	for _, v := range parts {
		t.Log(v)
	}
	t.Log(strings.Join(parts, "-"))

}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
