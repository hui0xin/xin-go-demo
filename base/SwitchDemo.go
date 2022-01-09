package main

import (
	"fmt"
	"time"
)

//SwitchDemo.go
func main() {

	//第一种常规
	num := 2
	fmt.Print("write ", num, " as : ")
	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//调用日期函数获取当前日期
	var date = time.Now().Weekday()
	fmt.Println(date)
	//第二种
	switch date {
	//可以使用逗号来分隔多个表达式。
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	fmt.Println("------------------------------")
	t := time.Now()
	fmt.Println(t)
	//不带表达式的 switch 是实现 if/else 逻辑的另一种方式。 这里还展示了 case 表达式也可以不使用常量。
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	fmt.Println("------------------------------")
	// 以借口方式
	//类型开关 (type switch) 比较类型而非值。可以用来发现一个接口值的类型。
	//在这个例子中，变量 t 在每个分支中会有相应的类型。
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
