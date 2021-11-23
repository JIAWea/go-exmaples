package main

import (
	"fmt"
	"sort"
)

func sliceSort() {
	// 标准库 排序
	a := []int{50, 10, 20, 5, 8, 10}
	sort.Ints(a)
	fmt.Println(a) // [5 8 10 10 20 50]

	b := []string{"a", "d", "c", "b"}
	sort.Strings(b)
	fmt.Println(b) // [a b c d]
}
