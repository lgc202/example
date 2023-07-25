package arr_index

//给你一个长度为 n 的整数数组 nums ，其中 nums 的所有整数都在范围 [1, n] 内，且每个整数出现 一次 或 两次 。请你找出所有出现 两次
//的整数，并以数组形式返回。
//
// 你必须设计并实现一个时间复杂度为 O(n) 且仅使用常量额外空间的算法解决此问题。
//
//
//
// 示例 1：
//
//
//输入：nums = [4,3,2,7,8,2,3,1]
//输出：[2,3]
//
//
// 示例 2：
//
//
//输入：nums = [1,1,2]
//输出：[1]
//
//
// 示例 3：
//
//
//输入：nums = [1]
//输出：[]
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 10⁵
// 1 <= nums[i] <= n
// nums 中的每个元素出现 一次 或 两次
//
//
// Related Topics 数组 哈希表 👍 726 👎 0

// 方法一: 变负数
func findDuplicates(nums []int) []int {
	var res []int
	for i := range nums {
		index := abs(nums[i]) - 1
		// 如果 index 处已经是负数, 说明已经去修改过一次了
		if nums[index] < 0 {
			res = append(res, index+1)
		} else {
			nums[index] = -nums[index]
		}
	}

	return res
}

// 方法二: 加 n
func findDuplicates2(nums []int) []int {
	n := len(nums)
	for i := range nums {
		// 与没加 n 前的 num[i] - 1 等效
		index := (nums[i] - 1) % n
		nums[index] += n
	}

	var res []int
	for i := range nums {
		// 大于 2*n, 说明至少加了两次n
		if nums[i] > 2*n {
			res = append(res, i+1)
		}
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
