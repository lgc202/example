package array

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	// 对撞双指针, 当发现和比目标大时, right 左移, 反之left右移
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}
