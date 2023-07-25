package prefix_sum

//给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
//
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在 32 位 整数范围内。
//
// 请不要使用除法，且在 O(n) 时间复杂度内完成此题。
//
//
//
// 示例 1:
//
//
//输入: nums = [1,2,3,4]
//输出: [24,12,8,6]
//
//
// 示例 2:
//
//
//输入: nums = [-1,1,0,-3,3]
//输出: [0,0,9,0,0]
//
//
//
//
// 提示：
//
//
// 2 <= nums.length <= 10⁵
// -30 <= nums[i] <= 30
// 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在 32 位 整数范围内
//
//
//
//
// 进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
//
// Related Topics 数组 前缀和 👍 1509 👎 0

// 前缀乘积
func productExceptSelf(nums []int) []int {
	// 时间复杂度 O(n^2), 会超时
	answer := make([]int, len(nums))
	for i := range nums {
		// 先求出左边的乘积
		leftProduct := 1
		for j := 0; j < i; j++ {
			leftProduct *= nums[j] // 有重复计算
		}

		// 再求出右边的乘积
		rightProduct := 1
		for j := i + 1; j < len(nums); j++ {
			rightProduct *= nums[j]
		}

		answer[i] = leftProduct * rightProduct
	}

	return answer
}

func productExceptSelf2(nums []int) []int {
	// 时间复杂度为 O(n), 空间复杂度为 O(2n) -> O(n)
	leftProduct := make([]int, len(nums))
	rightProduct := make([]int, len(nums))

	leftProduct[0] = 1
	for i := 1; i < len(nums); i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i-1] // 优化: 不需要重复计算
	}

	rightProduct[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		rightProduct[i] = rightProduct[i+1] * nums[i+1]
	}

	answer := make([]int, len(nums))
	for i := range nums {
		answer[i] = leftProduct[i] * rightProduct[i]
	}

	return answer
}

func productExceptSelf3(nums []int) []int {
	// 时间复杂度为 O(n), 空间复杂度为 O(1), 返回数组的空间不算在内

	// 先把左边乘积存到 answer 数组中
	answer := make([]int, len(nums))
	answer[0] = 1
	for i := 1; i < len(nums); i++ {
		answer[i] = answer[i-1] * nums[i-1] // 优化: 不需要重复计算
	}

	rightProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		// 利用左边乘积乘以右边乘积
		answer[i] = answer[i] * rightProduct
		// 更新右边乘积给下一次使用, 即计算 i-1 时使用
		rightProduct = rightProduct * nums[i]
	}

	return answer
}
