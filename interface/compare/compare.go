package main

import "fmt"

func main() {
	// 接口值比较有两种情形：
	// 1) 比较非接口值和接口值
	// 2) 比较两个接口值
	// 3) 一个[]T类型的值不能直接转换为[]I,即使T实现了接口类型I

	// 1) 将非接口值必须实现接口值的类型，所以先转为接口值，再直接比较两个接口值
	// 2) 判断两个接口值是否相等，需要满足下面的任何一种情况
	// 		- 两个接口值都为nil
	// 		- 两个接口值的动态类型相同，动态类型可为比较类型并且动态值相等
	var a, b, c interface{} = "abc", 123, "a" + "b" + "c"

	fmt.Println(a == c) // true
	fmt.Println(a == b) // false

	var x *int = nil
	var y *bool = nil
	var i interface{} = nil
	var ix, iy interface{} = x, y

	fmt.Println(ix == iy) // false
	fmt.Println(ix == i)  // false
	fmt.Println(iy == i)  // false

	// 3) []string值不能直接转换为[]interface
	words := []string{
		"go", "is", "a", "high", "efficient", "language",
	}

	// cannot use words (type []string) as type []interface {} in argument to fmt.Println
	// fmt.Println(words)	 // 编译成功 输出切片格式
	// fmt.Println(words...) // 编译失败

	iw := make([]interface{}, 0, len(words))

	for _, w := range words {
		iw = append(iw, w)
	}

	// iw 输出切片
	fmt.Println(iw) // [go is a high efficient language]
	// iw... 将切片输出为字符串
	fmt.Println(iw...) // go is a high efficient language

}
