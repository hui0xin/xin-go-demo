package main

import "fmt"

/**
这里是一个函数，接受两个 int 并且以 int 返回它们的和
Go 需要明确的 return，也就是说，它不会自动 return 最后一个表达式的值
 */
func plus(a int, b int) int {
	return a + b
}

//当多个连续的参数为同样类型时， 可以仅声明最后一个参数的类型，忽略之前相同类型参数的类型声明。
func plusPlus(a, b, c int) int {
	return a + b + c
}

//多返回值
//(int, int) 在这个函数中标志着这个函数返回 2 个 int。
func vals() (int, int) {
	return 3, 7
}

//可变参数的函数
//这个函数接受任意数量的 int 作为参数。
func sum(nums ...int) {
	//fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

//Go 支持匿名函数， 并能用其构造 闭包。
//匿名函数在你想定义一个不需要命名的内联函数时是很实用的。
//intSeq 函数返回一个在其函数体内定义的匿名函数。 返回的函数使用闭包的方式 隐藏 变量 i。 返回的函数 隐藏 变量 i 以形成闭包。
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//递归函数
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

//指针
//我们将通过两个函数：zeroval 和 zeroptr 来比较 指针 和 值。
//zeroval 有一个 int 型参数，所以使用值传递。
//zeroval 将从调用它的那个函数中得到一个实参的拷贝：ival。
func zeroval(ival int) {
	ival = 0
}
//zeroptr 有一个和上面不同的参数：*int，这意味着它使用了 int 指针。
//紧接着，函数体内的 *iptr 会 解引用 这个指针，从它的内存地址得到这个地址当前对应的值。 对解引用的指针赋值，会改变这个指针引用的真实地址的值。
func zeroptr(iptr *int) {
	*iptr = 0
}

//函数 demo
func main(){

	//通过 函数名(参数列表) 来调用函数，
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	var res2 = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res2)

	// 多返回值的调用
	a, b := vals()
	fmt.Println("a =",a,",b =",b)

	//如果你仅仅需要返回值的一部分的话，你可以使用空白标识符 _。
	_, c := vals()
	fmt.Println(c)

	fmt.Println("--------------------------")
	//变参函数使用常规的调用方式，传入独立的参数。
	sum(1, 2)
	sum(1, 2, 3)

	//如果你有一个含有多个值的 slice，想把它们作为参数使用， 你需要这样调用 func(slice...)。
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	fmt.Println("############################")
	//我们调用 intSeq 函数，将返回值（一个函数）赋给 nextInt。 这个函数的值包含了自己的值 i，这样在每次调用 nextInt 时，都会更新 i 的值。
	nextInt := intSeq()
	//调用多次，就可以执行多次内部函数
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	//为了确认这个状态对于这个特定的函数是唯一的，我们重新创建并测试一下
	newInts := intSeq()
	fmt.Println(newInts())

	//递归调用
	fmt.Println("*****************************")
	fmt.Println(fact(7))

	//指针调用
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	i := 1

	//值传递
	zeroval(i)
	fmt.Println("zeroval:", i)

	//通过 &i 语法来取得 i 的内存地址，即指向 i 的指针。
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)
}

