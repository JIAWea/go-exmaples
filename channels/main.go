package main

import (
	"fmt"
	"time"
)

func main() {

	// 无缓冲channel
	counter := make(chan int)

	// 有缓存channel
	nums := make(chan int, 3)

	// 给counter发送至并关闭
	go func() {
		time.Sleep(2 * time.Second)
		counter <- 1
		close(counter)
	}()

	fmt.Println("before print counter, need to wating 2 second...")
	fmt.Println("counter: ", <-counter)
	val, ok := <-counter
	if ok {
		fmt.Println("counter is not closed, val: ", val)
	}
	fmt.Println("after print counter.")

	go func() {
		nums <- 1
		nums <- 2
		nums <- 3
	}()

	for i := 0; i < cap(nums); i++ {
		fmt.Println(<-nums)
	}

	close(nums)

	// 使用for-range循环通道
	for v := range counter {
		fmt.Println(v)
	}
	// 上面for-range等价于
	for {
		x, xok := <-counter
		if !xok {
			break
		}
		fmt.Println(x)

	}
}
