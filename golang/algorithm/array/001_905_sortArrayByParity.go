package array

// https://leetcode.cn/problems/sort-array-by-parity/
// 给你一个整数数组 nums，将 nums 中的的所有偶数元素移动到数组的前面，后跟所有奇数元素。
// 返回满足此条件的 任一数组 作为答案。

// 示例 1：
// 输入：nums = [3,1,2,4]
// 输出：[2,4,3,1]
// 解释：[4,2,3,1]、[2,4,1,3] 和 [4,2,1,3] 也会被视作正确答案。

// 示例 2：
// 输入：nums = [0]
// 输出：[0]

// 思路: 使用对撞双指针, 左边指针从第零个元素开始, 右边指针从最后一个元素开始
// 1. 当左边元素是奇数并且右边元素是偶数时, 交换两个元素的位置
// 2. 当左边元素已经是偶数并且右边元素已经是奇数时, 只需要将左边的角标右移并且右边的角标左移
// 3. 当左边元素是奇数并且右边元素也是奇数时, 只需要右边的角标左移
// 4. 当左边元素是偶数并且右边元素也是偶数时, 只需要左边的角标右移

func sortArrayByParity(nums []int) []int {
	left := 0
	right := len(nums) - 1
	// 对撞双指针, 奇数放后面
	for left < right {
		if nums[left]%2 == 1 && nums[right]%2 == 0 {
			nums[left], nums[right] = nums[right], nums[left]
		} else if nums[left]%2 == 0 && nums[right]%2 == 1 {
			left++
			right--
		} else if nums[left]%2 == 1 && nums[right]%2 == 1 {
			right--
		} else {
			left++
		}
	}

	return nums
}
