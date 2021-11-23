package main

import "fmt"

func main() {
	// Go 中所有函数参数传递都是值拷贝。
	// 但参数是map, slice, chan等类型时，在函数中修改参数的值也是会影响源参数的内容，
	// 因为它们是引用类型，存储的是值的地址，所以当函数修改参数时，自然也就修改了底层的值
	s := []string{"1", "2", "3"}
	fmt.Printf("调用前指针: %p, s[1]指针: %p\n", &s, &s[1])

	test2(s)
	fmt.Println("s: ", s) // [1 6 3]

	m := map[string]string{"name": "ray", "age": "18"}
	modify(m)
	fmt.Println("m: ", m) // map[age:20 name:ray]
}

func test2(s []string) {
	s[1] = "6"
	fmt.Printf("调用后指针: %p, s[1]指针: %p\n", &s, &s[1])
}

func modify(m map[string]string) {
	m["age"] = "20"
}
