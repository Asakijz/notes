package lesson09

import (
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Logf("初始化的 string 值为： %#v", s) // 初始化为默认零值
	s = "hello"
	t.Logf("声明为 hello 后的长度是：%#v", len(s))
	//s[1] = '3' //string 是不可变的 byte slice
	s = "\xE4\xB8\xA5" // 可以存储任何二进制数据
	//s = "\xE4\xBA\xB5\xFF" // 不是正确的编码，但是依旧可以执行，len 取值同样以 byte 为准
	t.Logf("存储的二进制数据为：%v", s)
	t.Logf("存储的二进制数据的长度为：%v", len(s)) // 长度为 3 是因为上面是 3 个 byte
	s = "中"
	t.Logf("'中'这个字符串的长度为：%v", len(s))

	c := []rune(s)
	t.Logf("rune 的长度为：%v", len(c))
	//t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF-8 %x", s)
}

func TestStringToRune(t *testing.T) {
	// 遍历 rune
	s := "学习编程针不戳"
	for _, c := range s {
		// range 和字符串结合时，迭代输出的是 rune 而不是 byte
		t.Logf("%[1]c %[1]d", c)
		/* 正常后面需要跟两个参数，一个对应 %c，一个对应 %d
		但[1]代表这边都是跟第一个参数对应
		*/
	}
}
