package maptest_test

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := map[string]int{"key1": 1, "key2": 2}
	t.Log(m["key1"], len(m))

}

func TestAccess(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[0])
	m1[2] = 0
	t.Log(m1[2])
	//m1[3] = 6
	if v, ok := m1[3]; ok {
		t.Logf("Key 3 value is %d", v)
	} else {
		t.Log("Key 3 is not exist")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		fmt.Printf("key: %d, value:%d", k, v)
		t.Log(k, v)
	}

}
