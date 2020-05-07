package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 1. 字符串是否有某个前缀或后缀
	x := "What a beautiful day"
	y1 := strings.HasPrefix(x, "What")
	fmt.Println(y1) // true

	fmt.Println(strings.HasSuffix(x, "day")) //true

	// 2. 字符或子串在字符串中出现的位置
	fmt.Println(strings.Index(x, "bea"))   // 7
	fmt.Println(strings.IndexAny(x, "ba")) // 2 返回任何一个字符存在的位置
	// 有三个对应的查找最后一次出现的位置
	// func LastIndex(s, sep string) int
	// func LastIndexAny(s, chars string) int
	// func LastIndexFunc(s string, f func(rune) bool) int

	// 3. 是否存在某个字符或子串(利用Index查看是否有存在字符串)
	fmt.Println(strings.Contains(x, "beaa")) //false

	// func Contains(s, substr string) bool {
	// 	return Index(s, substr) >= 0
	// }

	// 4 字符串分割为[]string
	// 4.1 Split 和 SplitAfter、 SplitN 和 SplitAfterN
	fmt.Printf("%q\n", strings.Split("foo,bar,baz", ","))      // ["foo" "bar" "baz"]
	fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ",")) // ["foo," "bar," "baz"]
	fmt.Printf("%q\n", strings.SplitN("foo,bar,baz", ",", 2))  // ["foo" "bar,baz"]

	// 4.2 Fields 和 FieldsFunc
	// 分割一个或多个空格
	fmt.Println(strings.Fields("  foo bar  baz   "))                      // ["foo" "bar" "baz"]
	fmt.Println(strings.FieldsFunc("  foo bar  baz   ", unicode.IsSpace)) // [foo bar baz]
}
