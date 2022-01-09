package main

import "fmt"

//VarDemo.go 变量
func main() {
	var a = "test"
	fmt.Println(a)

	//可以一次声明多个变量
	//var c,d int = 1,2
	var c, d = 1, 2
	fmt.Println(c, d)

	//Go 会自动推断已经有初始值的变量的类型。
	var m = true
	fmt.Println(m)

	//声明后却没有给出对应的初始值时，变量将会初始化为 零值 。
	//例如，int 的零值是 0。
	var e int
	fmt.Println(e)

	//var f = "test2"
	//var f string = "test2"
	f := "test2"
	fmt.Println(f)

}
