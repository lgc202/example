package array

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
