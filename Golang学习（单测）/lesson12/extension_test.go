package lesson12

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) Speakto(host string) {
	p.Speak()
	fmt.Println("", host)
}

type Dog struct {
	p *Pet
	//Pet // 匿名嵌套类型
}

func (d *Dog) Speak() {
	//d.p.Speak()
	fmt.Println("Wang")
}

func (d *Dog) Speakto(host string) {
	//d.p.Speakto(host)
	d.Speak()
	fmt.Println("", host)
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Speakto("Chao")
}
