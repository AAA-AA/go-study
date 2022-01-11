package polymtest_test

import (
	"fmt"
	"testing"
)

type Code string

type Progarammer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World~\")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.println(\"Hello World!\")"
}

func write(p Progarammer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	//goProg := new(GoProgrammer)
	goProg := &GoProgrammer{}
	javaProg := new(JavaProgrammer)
	write(goProg)
	write(javaProg)
}
