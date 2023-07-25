package double_pointer

func removeDuplicates80_1(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	slow := 2
	fast := 2
	for fast < len(nums) {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}

		fast++
	}

	return slow
}
