package main

import (
	"fmt"
)

func main() {

	str1 := "thisisatest"
	str2 := "test"

	fmt.Println(strStr(str1, str2))
}

/*
	给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置
	(从 0 开始)。如果不存在，则返回 -1。
*/
func strStr(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		if i+len(needle) > len(haystack) {
			return -1
		}

		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}
