package lesson13

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	//if i, ok := p.(int); ok {
	//	fmt.Println("Interger", i)
	//	return
	//}
	//if s, ok := p.(string); ok {
	//	fmt.Println("String", s)
	//	return
	//}
	//fmt.Println("Unknow Type")

	// 可以用 switch case 写法来简化操作步骤
	switch v := p.(type) {
	case int:
		fmt.Println("Interger", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknow Type")
	}

}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}
