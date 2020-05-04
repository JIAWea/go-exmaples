package main

import (
	"fmt"
)

func main() {
	var x [5]int
	x[0] = 0
	x[1] = 5
	fmt.Println("values of x:", x) // [0 5 0 0 0]

	x[2] = 10
	x[3] = 15
	x[4] = 20
	fmt.Println("values of x:", x) // [0 5 10 15 20]

	y := [5]int{25, 30, 35, 40}
	fmt.Println("values of y:", y)   // [25 30 35 40 0]
	fmt.Println("len of y:", len(y)) // 5

	// ... 表示初始化多少个值长度就多少
	z := [...]int{45, 50, 55, 60, 65, 70}
	fmt.Println("values of z:", z)      // [45 50 55 60 65 70]
	fmt.Println("length of z:", len(z)) // 6

	// 初始化指定位置
	langs := [3]string{0: "33", 2: "33"}
	fmt.Println("values of langs:", langs) // [33  33]

	// ========== 循环数组 ==========
	// 获取索引和值
	for i, v := range langs {
		fmt.Printf("index: %v, value: %s\n", i, v)
	}

	// 只获取索引
	for k := range langs {
		fmt.Printf("index: %v\n", k)
	}
}
