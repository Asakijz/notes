package lesson07_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))
	// 没有初始值的声明
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))
	m3 := make(map[int]int, 10) //用make声明map，第二个参数声明的是map的cap（容量）
	t.Logf("len m3=%d", len(m3))
	//t.Log(cap(m3)) //cap函数不能用于map求值
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	// 在 map 中，在访问的 key 不存在时，仍会返回零值，不能通过返回 nil 来判断元素是否存在
	// 那么这种情况下我们要如何判断这个 value 是不存在呢还是本身就是零值？
	m1[2] = 0
	t.Log(m1[2])

	//m1[3] = 0
	// v 代表 map 的 value 值，ok 返回的是一个布尔值，如果该值不存在则为 false，反之为 true
	if v, ok := m1[3]; ok {
		t.Logf("key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing")
	}

}

func TestTravelMap(t *testing.T) {
	// 遍历循环取值
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
