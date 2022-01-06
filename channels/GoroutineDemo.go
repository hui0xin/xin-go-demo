package main

import (
	"fmt"
	"time"
)


func f(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}

//协程(goroutine) 是轻量级的执行线程。
func main(){

	//假设我们有一个函数叫做 f(s)。 我们一般会这样 同步地 调用它
	f("direct")

	//使用 go f(s) 在一个协程中调用这个函数。 这个新的 Go 协程将会 并发地 执行这个函数。
	go f("goroutine")

	//你也可以为匿名函数启动一个协程。
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//现在两个协程在独立的协程中 异步地 运行， 然后等待两个协程完成（更好的方法是使用 WaitGroup）。
	time.Sleep(time.Second)
	fmt.Println("done")
}


//当我们运行这个程序时，首先会看到阻塞式调用的输出，然后是两个协程的交替输出。
//这种交替的情况表示 Go runtime 是以并发的方式运行协程的。
