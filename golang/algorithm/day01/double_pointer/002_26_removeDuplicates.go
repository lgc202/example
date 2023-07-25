package double_pointer

// 删除有序数组中的重复项
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array
//
// 给你一个升序排列的数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持一致 。然后返回nums中唯一元素的个数。
//
// 考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
//
// 更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums的其余元素与 nums 的大小不重要。
// 返回 k。
// 判题标准:
//
// 系统会用下面的代码来测试你的题解:
// int[] nums = [...]; // 输入数组
// int[] expectedNums = [...]; // 长度正确的期望答案
//
// int k = removeDuplicates(nums); // 调用
//
// assert k == expectedNums.length;
// for (int i = 0; i < k; i++) {
//     assert nums[i] == expectedNums[i];
// }
//
// 示例 1：
//
// 输入：nums = [1,1,2]
// 输出：2, nums = [1,2,_]
// 解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2。不需要考虑数组中超出新长度后面的元素。

// 思路：使用快慢双指针
// slow 指向 已处理区域的下一个元素
// 刚开始时 0 号元素表示已经处理过了, 所以 slow 指向的是它的下一个位置, 也就是 slow = 1

func removeDuplicates26_1(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	slow := 1
	fast := 1
	for fast < len(nums) {
		if nums[slow-1] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}

		fast++
	}

	return slow
}

// 思路：使用快慢双指针
// slow 指向 已处理区域的最后一个元素
// 刚开始时 0 号元素表示已经处理过了, 所以 slow 指向它, 也就是 slow = 0

func removeDuplicates26_2(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	slow := 0
	fast := 1
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}

		fast++
	}

	return slow + 1
}
