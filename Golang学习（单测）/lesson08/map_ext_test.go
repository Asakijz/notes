package lesson08

import "testing"

func TestMapWithFunValue(t *testing.T) {
	// map 的 value 可以是一个函数
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

/* Go 的内置集合中没有 set 实现，可以 map[type]bool
   1.元素的唯一性
   2.基本操作
		1)添加元素
		2)判断元素是否存在
		3)删除元素
		4)元素个数

*/

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	// 直接把元素放在 key 的位置
	mySet[1] = true
	// 可以直接通过 value 来判断这个 key 是否存在
	n := 3
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	mySet[3] = true
	// 获取元素个数
	t.Log(len(mySet))
	// 删除元素
	delete(mySet, 1)
	// 这里图方便就把上面判断直接 copy 一份，正常应该抽象成一个函数来调用
	n = 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}
