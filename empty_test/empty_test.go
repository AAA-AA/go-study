package emptytest_test

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Integer", i)
	// 	return
	// }

	// if i, ok := p.(string); ok {
	// 	fmt.Println("string", i)
	// 	return
	// }
	// fmt.Println("Unknow Type")
	switch v := p.(type) {
	case int:
		fmt.Println("integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknow Type", v)
	}
}

func TestEmptyInterface(t *testing.T) {
	DoSomething(1)
	DoSomething("1")
}
