package main

import (
	"fmt"
	_ "fmt"
	"os"
)

/**
1. package main
2. func main()
3. main() 不能有返回值 os.Exit(int code)
4. 无法通过过 main() 传递参数
*/
func main() {
	if len(os.Args) > 1 {
		fmt.Print("Hello World", os.Args[1])
	}
	os.Exit(1)
}
