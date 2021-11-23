package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func doSomething(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("do something...")
	time.Sleep(time.Second * 2)
}

func doSomethingChan(ch chan<- struct{}) {
	fmt.Println("do something...")
	time.Sleep(time.Second * 2)
	ch <- struct{}{}
}

func main() {
	// var wg sync.WaitGroup

	// wg.Add(1)
	// go doSomething(&wg)

	// wg.Wait()
	// fmt.Println("1. done!")

	// ch := make(chan struct{})
	// go doSomethingChan(ch)
	// <-ch
	// fmt.Println("2. done!")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// defer cancel()

	go handle(ctx, 1*time.Second)

	cancel()
	select {
	case <-ctx.Done():
		fmt.Println("main, err: ", ctx.Err())
	}
}

func handle(ctx context.Context, t time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("[handle] wrong ", ctx.Err())
	case <-time.After(t):
		fmt.Println("[handle] timeout with ", t)
	}
}
