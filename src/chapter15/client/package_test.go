package client

import (
	"chapter15/series"
	"fmt"
	"testing"
)

func init() {
	fmt.Println("hello first")
}

func init() {
	fmt.Println("hello second")
}

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSeries(5))
}
