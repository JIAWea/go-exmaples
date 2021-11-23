// 这个示例程序展示goroutine调度器是如何在单个线程上
// 切分时间片的
package main

import (
	"fmt"
	"runtime"
	"sync"
)

//   wg用来等待程序完成
var wg sync.WaitGroup

func main() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	fmt.Println("Create Goroutines")

	go printPrime("A")
	go printPrime("B")
	// go func() {
	// 	defer wg.Done()

	// 	for i := 0; i < 10; i++ {
	// 		for char := 'a'; char < 'a'+26; char++ {
	// 			fmt.Printf("%c ", char)
	// 		}
	// 	}
	// }()

	// go func() {
	// 	defer wg.Done()

	// 	for i := 0; i < 10; i++ {
	// 		for char := 'A'; char < 'A'+26; char++ {
	// 			fmt.Printf("%c ", char)
	// 		}
	// 	}
	// }()

	fmt.Println("Waiting To   Finish")
	wg.Wait()
	fmt.Println("Terminating Program")
}

// printPrime 显示5000以内的素数值(大于1，除了1和本身外不能被整除)
func printPrime(prefix string) {
	defer wg.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}

	fmt.Println("Completed", prefix)
}
