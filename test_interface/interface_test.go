package interface_test

import (
	"testing"
)

type Progarammer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Progarammer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
