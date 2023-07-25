package arr_index

//给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数
//字，并以数组的形式返回结果。
//
//
//
// 示例 1：
//
//
//输入：nums = [4,3,2,7,8,2,3,1]
//输出：[5,6]
//
//
// 示例 2：
//
//
//输入：nums = [1,1]
//输出：[2]
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
//
//
// 进阶：你能在不使用额外空间且时间复杂度为 O(n) 的情况下解决这个问题吗? 你可以假定返回的数组不算在额外空间内。
//
// Related Topics 数组 哈希表 👍 1251 👎 0

// 方法一: 变负数
func findDisappearedNumbers(nums []int) []int {
	for i := range nums {
		index := abs(nums[i]) - 1
		// 一个数字只需要变一次即可, 否则负负得正
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}

	var res []int
	for i := range nums {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}

	return res
}

// 方法二: 加 n
func findDisappearedNumbers2(nums []int) []int {
	n := len(nums)
	for i := range nums {
		index := (nums[i] - 1) % n
		nums[index] += n
	}

	var res []int
	for i := range nums {
		// 注意是小于等于 n, 因为 num[i] 是从 1 开始的
		// 加上 n 后必定大于 n, 至少比 n 大 1
		if nums[i] <= n {
			res = append(res, i+1)
		}
	}

	return res
}
