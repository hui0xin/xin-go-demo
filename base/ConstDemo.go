package main

import (
	"fmt"
	"math"
)

//const 用于声明一个常量。
const s string = "我是一个常量"

// const 常量demo
func main(){

	fmt.Println(s)

	const i = 88
	fmt.Println(i)

	//运算
	const d = 3e20 / i
	fmt.Println(d)

	//自带函数
	fmt.Println(math.Sin(d))

}