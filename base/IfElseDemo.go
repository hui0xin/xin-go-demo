package main

import "fmt"

//IfElseDemo.go
func main() {

	//if else
	if 7%2 == 0 {
		fmt.Println("test1")
	} else {
		fmt.Println("test1")
	}

	//if -- else if----else
	//if num := 9; num < 0{
	num := 9
	//if (num < 0){
	if num < 0 {
		fmt.Println("if")
	} else if num < 10 {
		fmt.Println("else if")
	} else {
		fmt.Println("other")
	}

	//Go 没有三目运算符， 即使是基本的条件判断，依然需要使用完整的 if 语句。

}
