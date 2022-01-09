package main

import "fmt"

//ForDemo.go for循环
func main() {

	//第一种
	i := 1
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	//经典写法
	for j := 1; j < 10; j++ {
		fmt.Println(j)
	}

	//跳蛛循环
	for {
		fmt.Println("loop")
		break
		//return
	}
	fmt.Println("______________________________________")
	//跳出单次循环
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
