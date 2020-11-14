package map_ext

import "testing"

/*
	Map与工厂模式
		1. map的value可以是一个方法
		2. Go的Dock type 接口方式一起，可以方便实现单一方法对象的工厂模式
*/
func TestMapWithFunValue(t *testing.T) {
	m1 := map[int]func(op int) int{}
	m1[1] = func(op int) int {
		return op
	}

	m1[2] = func(op int) int {
		return op * op
	}

	m1[3] = func(op int) int {
		return op * op * op
	}

	for key, v := range m1 {
		t.Log(key, v(key))
	}
}

/*
	Go的内置集合中没有Set实现，可以map[type]bool
		set操作：
			1. 添加元素
			2. 判断一个元素是否存在
			3. 删除元素
			4. 元素个数
*/
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	t.Log(set(1, mySet))
	mySet[2] = true
	mySet[3] = true
	delete(mySet, 1)
	t.Log(set(1, mySet))
}

func set(key int, testMap map[int]bool) bool {
	if testMap[key] {
		return true
	} else {
		return false
	}
}
