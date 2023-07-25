package double_pointer

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
//
//
//
// 示例 1:
//
//
//输入: nums = [0,1,0,3,12]
//输出: [1,3,12,0,0]
//
//
// 示例 2:
//
//
//输入: nums = [0]
//输出: [0]
//
//
//
// 提示:
//
//
//
// 1 <= nums.length <= 10⁴
// -2³¹ <= nums[i] <= 2³¹ - 1
//
//
//
//
// 进阶：你能尽量减少完成的操作次数吗？
//

// 思路: 使用快慢指针 slow 和 fast
// 快慢指针会将数组分为三个区域
// (1) slow 之前的区域称为已处理并且满足题目要求区域
// (2) slow 和 fast 之间的区域称为已处理但不满足题目要求区域
// (3) fast 后面的是未处理的区域
// slow 可能指向的是已处理区域的最后一个元素, 也可能指向的是已处理区域的下一个元素

// slow 指向已处理区域的最后一个元素
// nums = [0,1,0,3,12]
//  1. 刚开始由于还没有处理任何元素 令 slow = -1, fast = 0
//  2. 当 fast = 1 时, 该位置上的元素为1, 不等于0, 需要交换 slow 和 fast 的位置。交换之前要让 slow++
//     最后得到 nums = [1,0,0,3,12], 此时 slow 指向的是 index 为 0 的区域, 为已处理区域的最后一个元素
func moveZeroes(nums []int) {
	slow := -1
	fast := 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			// 由于 slow 指向的是已经处理元素的最后一个
			// 所以要先将 slow 加 1 再赋值
			slow++
			nums[slow], nums[fast] = nums[fast], nums[slow] // 此时 slow 处非零, 属于已处理区域的最后一个元素
		}

		fast++
	}
}

// slow 指向已处理区域的下一个元素
// nums = [0,1,0,3,12]
//  1. 刚开始令 slow = 0, fast = 0
//  2. 当 fast = 1 时, 该位置上的元素为1, 不等于0, 需要交换 slow 和 fast 的位置。交换之后要让 slow++ 指向已处理区域的下一个
//     最后得到 nums = [1,0,0,3,12], 此时 slow 指向的是 index 为 1 的区域, 为已处理区域的下一个位置
func moveZeroes2(nums []int) {
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			// 可以进行优化的点: 当 slow 和 fast 指向同一个元素时, 不需要自己跟自己交换
			if slow != fast {
				// 由于 slow 指向的是已经处理元素的下一个
				// 所以要先赋值再将 slow 加 1
				nums[slow], nums[fast] = nums[fast], nums[slow]
			}
			slow++ // 此时 slow 指向已处理区域的下一个元素
		}

		fast++
	}
}

// 还可以继续优化: 不交换 slow 和 fast 的值, 而是通过直接把fast处的值直接赋值给slow处的方式
// 这种方法会比交互两个数快, 但要注意fast走完后要将slow后面的元素清0
func moveZeroes3(nums []int) {
	// slow 指向的是已处理元素的下一个
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			// 由于 slow 指向的是已经处理元素的下一个
			// 所以要先赋值再将 slow 加 1
			nums[slow] = nums[fast]
			slow++
		}

		fast++
	}

	// 最后把 slow 后面的数据清 0, 要注意当前 slow 指向的是已处理元素的下一个
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeroes4(nums []int) {
	// slow 指向的是已处理元素的最后一个
	slow := -1
	fast := 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			// 由于 slow 指向的是已经处理元素的最后一个
			// 所以要先将 slow 加 1 再赋值
			slow++
			nums[slow] = nums[fast] // 此时 slow 处非零, 属于已处理区域的最后一个元素
		}

		fast++
	}

	// 最后把 slow 后面的数据清 0, 要注意当前 slow 指向的是已处理元素的最后一个
	for i := slow + 1; i < len(nums); i++ {
		nums[i] = 0
	}
}
