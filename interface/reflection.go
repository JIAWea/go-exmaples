package main

import "fmt"

// Writer interface
type Writer interface {
	Write(buf []byte) (int, error)
}

// DocumentWriter struct
type DocumentWriter struct{}

func (DocumentWriter) Write(buf []byte) (int, error) {
	return len(buf), nil
}

func main() {
	// 接口值反射
	// 1) 断言
	// 2) type-switch

	// =================================================================
	// 1.1 断言基本类型
	// 编译器自动把123的类型推断为内置类型int
	var x interface{} = 123
	n, ok := x.(int)
	fmt.Println(n, ok) // 123 true

	// a := x.(float64) // 不好的写法，如果不是此类型的话则会引起painc
	a, ok := x.(float64)
	fmt.Println(a, ok) // 0 false

	// 1.2 断言接口类型
	var w Writer
	var k bool
	var i interface{} = DocumentWriter{}
	var y interface{} = "abc"

	w, k = i.(Writer)
	fmt.Println(w, k) // {} true

	x, k = w.(interface{}) // 任何类型都实现了空接口
	fmt.Println(x, k)      // {} true

	y, k = y.(Writer) // y的动态类型是string，且没实现Writer接口
	fmt.Println(y, k) // nil false

	// =================================================================
	// 2 type-switch流程控制代码块
	values := []interface{}{
		456, "abc", true, 0.33, int32(789), []int{1, 2, 3}, map[string]bool{}, nil,
	}

	for _, x := range values {
		// 如果不关心断言结果v，则可以写成 switch x.(type) {}
		switch v := x.(type) {
		case []int:
			fmt.Println("int slice: ", v)
		case string:
			fmt.Println("string: ", v)
		case int, float64, int32:
			fmt.Println("number: ", v)
		case nil:
			fmt.Println(v)
		default:
			fmt.Println("others:", v)
		}
	}
}
