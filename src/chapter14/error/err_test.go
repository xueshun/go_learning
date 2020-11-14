package error

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var LessThanTwoError error = errors.New("n must be not less than 2")
var GreaterThenHundredError error = errors.New("n shou be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, GreaterThenHundredError
	}
	fibList := []int{1, 1}
	for i := 0; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(1); err != nil {
		if err == LessThanTwoError {
			fmt.Println("It is less")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func GetFibonacci1(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println("Error", err)
	}
}

func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(list)
}

func TestGetFibonacci1(t *testing.T) {
	GetFibonacci2("1")
}
