package main

import "fmt"

func sliceStack() {
	// 1. 通过切片模拟栈
	stack := make([]int, 0)
	stack = append(stack, 10, 20, 30)

	// 出栈
	v := stack[len(stack)-1] // 30
	stack = stack[:len(stack)-1]
	fmt.Println("pop value: ", v, "\nnew stack: ", stack)

	// 进栈
	stackCopy := make([]int, len(stack))
	num := copy(stackCopy, stack)
	fmt.Println(num)
}
