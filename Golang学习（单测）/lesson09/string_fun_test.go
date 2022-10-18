package lesson09

import (
	"strconv"
	"strings"
	"testing"
)

// 常用字符串函数
func TestStringFn(t *testing.T) {
	s := "A,B,C"
	// 分割字符串
	parts := strings.Split(s, ",") // 第一个参数为需要分割的字符串，第二个参数为分割符
	for _, v := range parts {
		t.Log(v)
	}
	// 连接字符串
	t.Log(strings.Join(parts, "-")) // 第一个参数为需要连接的字符串，第二个参数为连接符
}

func TestStringConv(t *testing.T) {
	// 将整形转换为字符串
	// Itoa = Int to ASCII? 个人理解
	s := strconv.Itoa(10)
	t.Log("str" + s)
	// Atoi = ASCII to Int? 个人理解
	// Atoi 返回两个参数，第一个为转换完的 int，第二个为转换失败的 err
	if i, err := strconv.Atoi("10"); err != nil {
		t.Log("the argument is wrong,please check!")
	} else {
		t.Log(10 + i)
	}

}
