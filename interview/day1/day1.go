package main

import (
	"fmt"
	"sync"
)

var (
	wg  sync.WaitGroup
	ch1 = make(chan struct{})
	ch2 = make(chan struct{})
)

// 启动两个线程, 一个输出 1,3,5,7…99, 另一个输出 2,4,6,8…100 最后 STDOUT 中按
func main() {
	wg.Add(2)

	go thread1()
	go thread2()

	wg.Wait()
	fmt.Println("done")
}

func thread1() {
	for i := 1; i < 100; i += 2 {
		ch1 <- struct{}{}
		fmt.Println("thread1: ", i)
		<-ch2
	}

	wg.Done()
}

func thread2() {
	for i := 2; i < 101; i += 2 {
		<-ch1
		fmt.Println("thread2: ", i)

		ch2 <- struct{}{}
	}
	wg.Done()
}
