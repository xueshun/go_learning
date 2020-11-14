package _map

import "testing"

/*
	测试Map初始化
*/
func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	// 注意不能使用cap()访问map容量
	t.Logf("len m1=%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))

	//make(map[int]int, 10) 10 Initial Capacity
	// 为什么不初始化len?
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

/*
	测试Map中不存的Key
		前提描述
			1： map[int]int{}, 初始化为0
			2： 如果一个key的value为0， 那该怎么判断
*/
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])

	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("Key 3`s value is %d", v)
	} else {
		t.Log("Key 3 is not exiting")
	}
}

/*
	遍历Map
*/
func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}

	m2 := map[int]string{1: "hello", 2: "World"}
	for k, v := range m2 {
		t.Log(k, v)
	}
}
