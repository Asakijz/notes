package lesson3

import "testing"

type MyInt int64

//数据类型
//测试隐式转换
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	//即使自定义类型与对应变量类型相匹配也不行
	//c = b
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//在Go中无法对指针进行运算
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	//%T展示对应类型
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	s = "hello"
	//string类型默认值为空字符串
	t.Log("*" + s + "*")
	t.Log(len(s))
	if len(s) > 0 {
		t.Log("s len is ", len(s))
	}
}
