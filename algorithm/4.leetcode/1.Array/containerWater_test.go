package leetcode

import "testing"

func Test_mostWater(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostWater(tt.args.nums); got != tt.want {
				t.Errorf("mostWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
