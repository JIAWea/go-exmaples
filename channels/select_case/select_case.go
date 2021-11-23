package main

import "fmt"

func main() {
	// select-case分支流程和switch-case相似，有若干个case和最多一个default分支
	// 但也有许多不同点：
	// 1) select 和 { 之间不能有表达式
	// 2) 每个case必须跟着一个通道接受数据或者发送数据的操作
	// 3) 所有非阻塞的case操作中将有一个被随机选择执行（而不是从上到下的顺序）
	// 4) 在所有的case操作均为阻塞的情况下，如果default存在，则执行default，否则当前协程将进入阻塞状态
	// 5) 一个不含任何分支的select-case代码块将被当前协程处于永久阻塞状态

	var c chan struct{}

	// 若c是一个有缓冲的通道，然后向c发送数据，如果缓冲已满的话，则会执行default分支
	select {
	case <-c:

	case c <- struct{}{}:

	default:
		fmt.Println("Go here.")
	}

	b := make(chan int)
	b <- 1
	select {
	case b <- 2:
		fmt.Println(b)
	case <-b:
		fmt.Println("recevie b")
	default:
		fmt.Println("default")
	}

}
