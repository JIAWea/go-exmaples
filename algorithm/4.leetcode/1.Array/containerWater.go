package leetcode

// https://leetcode-cn.com/problems/container-with-most-water/

/*

题目：
	给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，
垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以
容纳最多的水。

示例：
	输入：[1,8,6,2,5,4,8,3,7]
	输出：49

思路：
	通过指针碰撞，首尾 2 个指针，每次移动以后都判断长宽的乘积是否最大
*/
func mostWater(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {

	}

	return max
}
