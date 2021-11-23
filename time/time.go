package main

import (
	"fmt"
	"time"
)

func main() {
	// t1 := time.Now().UnixNano()
	// t2 := time.Now().Local().Unix()
	// fmt.Println(strconv.Itoa(int(t1)))
	// fmt.Println(t2)

	for i := 0; i < 50; i++ {
		fmt.Println(time.Now().Format("20060102150405"))
	}
}
