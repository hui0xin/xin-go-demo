package main

import (
	"fmt"
	"os"
)

// EExitDemo.go 退出
//使用 os.Exit 可以立即以给定的状态退出程序。
//注意，不像例如 C 语言，Go 不使用在 main 中返回一个整数来指明退出状态。 如果你想以非零状态退出，那么你就要使用 os.Exit。
func main() {

	//当使用 os.Exit 时 defer 将不会 被执行， 所以这里的 fmt.Println 将永远不会被调用。
	defer fmt.Println("!")

	//退出并且退出状态为 3。
	os.Exit(3)
}
