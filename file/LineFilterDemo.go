package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//LineFilterDemo.go  行过滤器
//它读取 stdin 上的输入，对其进行处理，然后将处理结果打印到 stdout。 grep 和 sed 就是常见的行过滤器。
func main() {

	//用带缓冲的 scanner 包装无缓冲的 os.Stdin，
	//这为我们提供了一种方便的 Scan 方法， 将 scanner 前进到下一个 令牌（默认为：下一行）。
	scanner := bufio.NewScanner(os.Stdin)

	//Text 返回当前的 token，这里指的是输入的下一行。
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())

		//输出转换为大写后的行。
		fmt.Println(ucl)
	}

	//检查 Scan 的错误。 文件结束符（EOF）是可以接受的，它不会被 Scan 当作一个错误。
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

//试一下我们的行过滤器，首先创建一个有小写行的文件。
//$ echo 'hello'   > /tmp/lines
//$ echo 'filter' >> /tmp/lines

//然后使用行过滤器来得到大写的行。
//$ cat /tmp/lines | go run line-filters.go
//HELLO
//FILTER
