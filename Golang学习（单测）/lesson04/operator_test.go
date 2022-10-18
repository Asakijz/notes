package lesson04

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	//长度一致元素不相等比较为false
	b := [...]int{1, 3, 4, 5}
	t.Log(a == b)
	//长度不同的数组是无法做比较的，无法通过编译
	//c := [...]int{1, 2, 3, 4, 5}
	//t.Log(a == c)
	d := [...]int{1, 2, 3, 4}
	t.Log(a == d)

}

//按位置进行清零运算-类似Linux上的0644，0755权限
//&^ 按位置零-只要右边的二进制位为1，则结果为0
//1 &^ 0 -- 1
//1 &^ 1 -- 0
//0 &^ 1 -- 0
//0 &^ 0 -- 0

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestStatus(t *testing.T) {
	a := 7 //0111
	t.Log(Readable, Writeable, Executable)
	//这一步相等于 0111 &^ 1 -- 0
	a = a &^ Readable
	a = a &^ Writeable
	a = a &^ Executable
	t.Log(a&Readable == Readable)
	t.Log(a&Writeable == Writeable)
	t.Log(a&Executable == Executable)
}
