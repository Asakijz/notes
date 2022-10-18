package _const

import "testing"

//const定义常量
//iota为自增，初始为0
//定义一周七天，这里偷懒写3天
const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

func TestConstDay(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

//定义读写和执行，通过位运算判断状态
const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestStatus(t *testing.T) {
	a := 7 //0111
	//通过与运算来判断，就跟Linux上的0644，0755一样
	t.Log(a&Readable == Readable)
	t.Log(a&Writeable == Writeable)
	t.Log(a&Executable == Executable)
}
