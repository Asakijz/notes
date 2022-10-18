package exchange

import "testing"

func TestExchange(t *testing.T) {
	a := 1
	b := 2

	//其他语言默认需要一个中间量来进行两数交换
	//tmp := a
	//a = b
	//b = tmp

	//在Go中只需要下列操作即可交换两数
	a, b = b, a
	t.Log(a, b)
}
