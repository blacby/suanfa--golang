package _长度最小子数组

// 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
//
// 示例：
//
// 输入：s = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
// 提示：
//
// 1 <= target <= 10^9
// 1 <= nums.length <= 10^5
// 1 <= nums[i] <= 10^5
// leetcode: https://leetcode.cn/problems/minimum-size-subarray-sum/description/

//此解法超出时间限制
//func minSubArrayLen(target int, nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	result, right := 0, 0
//	for i := 0; i < len(nums); i++ {
//		temp := nums[i]
//		if temp >= target {
//			return 1
//		}
//		right = i
//		for {
//			right++
//			if right <= len(nums)-1 {
//				temp += nums[right]
//				if temp >= target {
//					result_temp := right - i + 1
//					if result > 0 && result_temp < result {
//						result = result_temp
//					} else if result == 0 {
//						result = result_temp
//					}
//					break
//				}
//			} else {
//				break
//			}
//		}
//	}
//	return result
//}

//自己没看题解写的 ac的解法
//·func minSubArrayLen(target int, nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	result, right, flag := 0, 0, 0
//	for i := 0; i < len(nums); i++ {
//		temp := nums[i]
//		if temp >= target {
//			return 1
//		}
//		right = i
//		for {
//			right++
//			if right <= len(nums)-1 {
//				temp += nums[right]
//				if temp >= target {
//					flag = 1
//					result_temp := right - i + 1
//					if result > 0 && result_temp < result {
//						result = result_temp
//					} else if result == 0 {
//						result = result_temp
//					}
//					break
//				}
//			} else {
//				break
//			}
//		}
//		//尝试修改
//		if right >= len(nums)-1 && flag != 1 {
//			break
//		}
//		flag = 0
//	}
//	return result
//}

//// 看了题解，自己写出的写法
//func minSubArrayLen(target int, nums []int) int {
//	result := len(nums) + 1
//	sum := 0
//
//	for i := 0; i < len(nums); i++ {
//		sum += nums[i]
//		j := 0
//		tempsum := sum
//		for tempsum >= target {
//			rempResult := i - j + 1
//			if rempResult < result {
//				result = rempResult
//			}
//			j++
//			tempsum -= nums[j-1]
//		}
//	}
//
//	if result == len(nums)+1 {
//		return 0
//	} else {
//		return result
//	}
//}

// 题解写法（copy）
func minSubArrayLen(target int, nums []int) int {
	i := 0
	l := len(nums)  // 数组长度
	sum := 0        // 子数组之和
	result := l + 1 // 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
	for j := 0; j < l; j++ {
		sum += nums[j]
		for sum >= target {
			subLength := j - i + 1
			if subLength < result {
				result = subLength
			}
			sum -= nums[i]
			i++ //这个是滑动的精髓，省不少事，算完之后再滑动，即使退出了，下次 j 向右了，i也直接是王4左了一位 ，千万不要从 i=0开始，很容易这样误解。
		}
	}
	if result == l+1 {
		return 0
	} else {
		return result
	}
}
