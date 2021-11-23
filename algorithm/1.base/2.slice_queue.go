package main

import "fmt"

func sliceQueue() {
	// 2. 通过切片模拟队列
	queue := make([]int, 0)

	// 入队
	queue = append(queue, 10, 20, 30)

	// 出队
	v := queue[0] // 30
	queue = queue[1:]
	fmt.Println("pop value: ", v, "\nnew queue: ", queue)
}
