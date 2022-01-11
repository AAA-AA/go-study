package testloop_test

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}

func TestSwitch(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")

		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondtion(t *testing.T) {

}
