package main

import (
	"fmt"
	"os"
	"strings"
)

// EnvDemo.go 环境变量
//环境变量 是一种向 Unix 程序传递配置信息的常见方式。 让我们来看看如何设置、获取以及列出环境变量。
func main() {

	//使用 os.Setenv 来设置一个键值对。 使用 os.Getenv获取一个键对应的值。 如果键不存在，将会返回一个空字符串。
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	//使用 os.Environ 来列出所有环境变量键值对。 这个函数会返回一个 KEY=value 形式的字符串切片。 你可以使用 strings.SplitN 来得到键和值。这里我们打印所有的键。
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}

//运行这个程序，显示我们在程序中设置的 FOO 的值， 然而没有设置的 BAR 是空的。
//$ go run environment-variables.go

//键的列表是由你的电脑情况而定的。
//
//TERM_PROGRAM
//PATH
//SHELL
//...

//如果我们在运行前设置了 BAR 的值，那么运行程序将会获取到这个值。
//$ BAR=2 go run environment-variables.go
