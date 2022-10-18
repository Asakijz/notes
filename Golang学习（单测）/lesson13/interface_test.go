package lesson13

import (
	"fmt"
	"testing"
)

//type Programmer interface {
//	WriteHelloWorld() string
//}
//
//type GoProgrammer struct {
//}
//
//func (g *GoProgrammer) WriteHelloWorld() string {
//	return "fmt.Println(\"hello world\")"
//}
//
//func TestClient(t *testing.T) {
//	var p Programmer
//	p = new(GoProgrammer)
//	t.Log(p.WriteHelloWorld())
//}

type Code string
type Programmer interface {
	WriteHelloWorld() Code
}
type GoProgrammer struct {
}

func (g GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World!\")"
}

func WriteFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	WriteFirstProgram(goProg)
	WriteFirstProgram(javaProg)

}
