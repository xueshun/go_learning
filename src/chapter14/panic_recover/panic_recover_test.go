package panic_recover

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

/*
	os.Exit()， 不会执行defer
*/
func TestOsExitVxExit(t *testing.T) {
	defer func() {
		fmt.Println("Hello World")
	}()
	fmt.Println("Start")
	os.Exit(1)
}

/*
	panic 退出会执行 defer
*/
func TestPanicExit(t *testing.T) {
	defer func() {
		fmt.Println("Hello World")
	}()
	fmt.Println("Start")
	panic("eee")
	fmt.Println("End")
}

func TestPanicExit1(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from", err)
		}
	}()

	fmt.Println("Start")
	errors.New("Something wrong!")
	panic("hei hei")
	fmt.Println("End")
}
