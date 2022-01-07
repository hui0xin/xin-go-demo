package main

import (
	"fmt"
	"net"
	"net/url"
)

//UrlDemo.go
//Go 提供内建的 JSON 编码解码（序列化反序列化）支持， 包括内建及自定义类型与 JSON 数据之间的转化。
func main() {

	//我们将解析这个 URL 示例，它包含了一个 scheme、 认证信息、主机名、端口、路径、查询参数以及查询片段。
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	//解析这个 URL 并确保解析没有出错。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	//直接访问 scheme。
	fmt.Println(u.Scheme)

	//User 包含了所有的认证信息， 这里调用 Username 和 Password 来获取单独的值。
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	//Host 包含主机名和端口号（如果存在）。使用 SplitHostPort 提取它们。
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	//这里我们提取路径和 # 后面的查询片段信息。
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	//要得到字符串中的 k=v 这种格式的查询参数，可以使用 RawQuery 函数。 你也可以将查询参数解析为一个 map。
	//已解析的查询参数 map 以查询字符串为键， 已解析的查询参数会从字符串映射到到字符串的切片， 因此如果您只想要第一个值，则索引为 [0]。
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
