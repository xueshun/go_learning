package array_test

import "testing"

/*
 数组初始化
*/
func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr2)
}

/*
 多维数组声明
*/
func TestMultiArray(t *testing.T) {
	c := [2][2]int{{1, 2}, {3, 4}}
	t.Log(c)
}

/*
 数组遍历
*/
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4}
	for i := 0; i < len(arr3); i++ {
		t.Log("for 遍历 ", arr3[i])
	}

	for idx, e := range arr3 {
		t.Log(idx, e)
	}

	for _, e := range arr3 {
		t.Log(e)
	}
}

/*
  数组截取
*/
func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3Sec := arr3[:]
	t.Log(arr3[1:2])
	t.Log(arr3[1:3])
	t.Log(arr3[1:])
	t.Log(arr3[:3])
	t.Log(arr3Sec)
}
