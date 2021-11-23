package leetcode

// https://leetcode.com/problems/two-sum/

/*

题目：
	在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。

示例：
	Given nums = [2, 7, 11, 15], target = 9
	return [0, 1]

*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		tmp := target - nums[i]
		if _, ok := m[tmp]; ok {
			return []int{m[tmp], i}
		}
		m[nums[i]] = i
	}

	return nil
}
