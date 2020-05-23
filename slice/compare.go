package main

import (
	"fmt"
	"reflect"
)

func StringSliceLoopCompare(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if (s1 == nil) != (s2 == nil) {
		return false
	}

	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}

func StringSliceReflectEqual(s1, s2 []string) bool {
	return reflect.DeepEqual(s1, s2)
}

func main() {
	// 两种比较两个切片是否相等的方法
	s1 := []string{"hi", "this", "is", "a", "test", "for", "two", "slice"}
	s2 := []string{"hi", "this", "is", "a", "test", "for", "two", "slice"}
	s3 := []string{"hi", "this", "is", "a", "test", "for", "two", "haha"}

	ok1 := StringSliceLoopCompare(s1, s2)
	fmt.Println("Equal:", ok1) // true
	ok2 := StringSliceLoopCompare(s1, s3)
	fmt.Println("Equal:", ok2) // false

	ok3 := stringSliceReflectEqual(s1, s2)
	fmt.Println("Equal:", ok3) // true
	ok4 := stringSliceReflectEqual(s1, s3)
	fmt.Println("Equal:", ok4) // false
}
