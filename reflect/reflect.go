package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	// 1. 通过反射判断一个值的类型
	for _, v := range []interface{}{"hi", 123, true, 90.99, func() {}} {
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			fmt.Printf("%s 是一个字符串\n", v.String())
		case reflect.Bool:
			fmt.Printf("%t 是一个布尔值\n", v.Bool())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Printf("%d 是一个数字\n", v.Int())
		default:
			fmt.Printf("%s 其他类型\n", v.Kind())
		}
	}

	// 2. 判断一个struct中的字段标签
	type S struct {
		F string `color:"red"`
	}
	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color")) // red

	// 3. 判断两个对象是否相等
	// ps: 使用DeepEqual的性能比自己遍历循环还差
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	fmt.Println(reflect.DeepEqual(a, b)) // true

	// 4. 判断一个类型是否实现了该接口
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	fileType := reflect.TypeOf((*os.File)(nil))
	fmt.Println(fileType.Implements(writerType)) // true

	// =================== 断言 ============================
	for _, v := range []interface{}{"hi", 123, true, 90.99, func() {}} {
		switch v.(type) {
		case string:
			fmt.Println("字符串:", v)
		case int:
			fmt.Println("数字:", v)
		case bool:
			fmt.Println("布尔值:", v)
		default:
			fmt.Println("其他类型:", v)
		}
	}
}
