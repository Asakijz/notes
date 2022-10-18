package lesson1

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20} // 可以不指定名称赋值，不过赋值顺序按照定义结构体的顺序
	e1 := Employee{
		Name: "Mike",
		Age:  30,
	}
	e2 := new(Employee) // 这种声明的是指针类型
	e2.Id = "2"
	e2.Name = "Rose"
	e2.Age = 22
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
}

// 所有数据的存放都在同一块位置上
//func (e *Employee) String() string {
//	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
//}

// 这种写法是相当于复制了一份，会有更大的空间复制开销
func (e Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	//e := &Employee{"0", "Bob", 20}
	t.Log(e.String())
}
