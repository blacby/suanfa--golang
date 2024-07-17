package _有序数组的平方

// 力扣题目链接:https://leetcode.cn/problems/squares-of-a-sorted-array/description/
//
// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//
// 示例 1：
//
// 输入：nums = [-4,-1,0,3,10]
// 输出：[0,1,9,16,100]
// 解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]
// 示例 2：
//
// 输入：nums = [-7,-3,2,3,11]
// 输出：[4,9,9,49,121]
func sortedSquares(nums []int) []int {

	if len(nums) < 1 {
		return []int{}
	}
	var left = 0
	var right = len(nums) - 1
	tempArray := make([]int, len(nums))
	var index = len(nums) - 1
	for {
		if right < left {
			return tempArray
		}
		if nums[left]*nums[left] < nums[right]*nums[right] {
			tempArray[index] = nums[right] * nums[right]
			index--
			right--
		} else {
			tempArray[index] = nums[left] * nums[left]
			index--
			left++
		}
	}
}
