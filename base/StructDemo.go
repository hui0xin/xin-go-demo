package main

import "fmt"

//这里的 person 结构体包含了 name 和 age 两个字段。
type person struct {
	name string
	age  int
}

//结构体 demo
func main(){

	//使用这个语法创建新的结构体元素。
	var parson1 = person{"Bob", 20}
	fmt.Println(parson1)

	//你可以在初始化一个结构体元素时指定字段名字。
	var parson2 = person{name: "Alice", age: 30}
	fmt.Println(parson2)

	//省略的字段将被初始化为零值。
	var person3 = person{name: "Fred"}
	fmt.Println(person3)

	//& 前缀生成一个结构体指针。
	var person4 = &person{name: "Ann", age: 40}
	fmt.Println(person4)

	//使用.来访问结构体字段。
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	//也可以对结构体指针使用. - 指针会被自动解引用。
	sp := &s
	fmt.Println(sp.age)

	//结构体是可变(mutable)的。
	sp.age = 51
	fmt.Println(sp.age)
}

