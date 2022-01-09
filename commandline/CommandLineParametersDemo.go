package main

import (
	"fmt"
	"os"
)

// CommandLineParametersDemo.go 命令行参数
// 命令行参数 是指定程序运行参数的一个常见方式。例如，go run hello.go， 程序 go 使用了 run 和 hello.go 两个参数。
//要实验命令行参数，最好先使用 go build 编译一个可执行二进制文件
func main() {

	//os.Args 提供原始命令行参数访问功能。 注意，切片中的第一个参数是该程序的路径， 而 os.Args[1:]保存了程序全部的参数。
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	//你可以使用标准的下标方式取得单个参数的值。
	arg := os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

//要实验命令行参数，最好先使用 go build 编译一个可执行二进制文件
//$ go build command-line-arguments.go

//$ ./command-line-arguments a b c d
//[./command-line-arguments a b c d]
//[a b c d]
//c
