package main

import (
	"fmt"
	"math"
)

//const 用于声明一个常量。
const s string = "我是一个常量"

// ConstDemo.go 常量
// Go 支持字符、字符串、布尔和数值 常量 。
func main() {

	fmt.Println(s)

	//const 语句可以出现在任何 var 语句可以出现的地方
	const i = 88
	fmt.Println(i)

	//常数表达式可以执行任意精度的运算
	const d = 3e20 / i
	fmt.Println(d)

	//自带函数
	fmt.Println(math.Sin(d))

}
