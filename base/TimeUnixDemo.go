package main

import (
	"fmt"
	"time"
)

//TimeUnixDemo.go 时间戳操作
//般程序会有获取 Unix 时间 的秒数，毫秒数，或者微秒数的需求。来看看如何用 Go 来实现。
func main() {

	//分别使用 time.Now 的 Unix 和 UnixNano， 来获取从 Unix 纪元起，到现在经过的秒数和纳秒数。
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	//注意 UnixMillis 是不存在的，所以要得到毫秒数的话， 你需要手动的从纳秒转化一下。
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	//你也可以将 Unix 纪元起的整数秒或者纳秒转化到相应的时间。
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
