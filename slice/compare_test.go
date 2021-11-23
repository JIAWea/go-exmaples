package main

import (
	"testing"
)

// goos: windows
// goarch: amd64
// BenchmarkStringSliceLoopCompare-4       30000000                36.4 ns/op
// BenchmarkStringSliceReflectEqual-4       1000000              1682 ns/op
// PASS
// ok      command-line-arguments  2.926s

var (
	s1 = []string{"hi", "this", "is", "a", "test", "for", "two", "slice"}
	s2 = []string{"hi", "this", "is", "a", "test", "for", "two", "slice"}
)

func BenchmarkStringSliceLoopCompare(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringSliceLoopCompare(s1, s2)
	}
}

func BenchmarkStringSliceReflectEqual(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringSliceReflectEqual(s1, s2)
	}
}
