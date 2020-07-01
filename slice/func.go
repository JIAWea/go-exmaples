package main

import "fmt"

func main() {
	// Go 中所有函数参数传递都是值拷贝。
	// 但参数是map, slice, chan等类型时，在函数中修改参数的值也是会影响源参数的内容，
	// 因为它们是引用类型，存储的是值的地址，所以当函数修改参数时，自然也就修改了底层的值

	s := []string{"1", "2", "3"}
	fmt.Printf("调用前指针: %p\n", &s)

	test1(s)
	test2(s)

	m := map[string]string{"name": "ray", "age": "18"}
	modify(m)
	fmt.Println("m: ", m) // map[age:20 name:ray]
}

func test1(s []string) {
	fmt.Printf("调用后指针: %p\n", &s)
}

func test2(s []string) {
	fmt.Printf("调用后指针: %p\n", &s)
}

func modify(m map[string]string) {
	m["age"] = "20"
}
