package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Print("wang")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Chao") //说明不支持继承
}
