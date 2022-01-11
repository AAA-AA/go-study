package client

import (
	"testing"
	"web-demo2/service"
)

func TestPackage(t *testing.T) {
	t.Log(service.GetFibonacci(5))
}
