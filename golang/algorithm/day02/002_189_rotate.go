package day02

//给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
//
//
//
// 示例 1:
//
//
//输入: nums = [1,2,3,4,5,6,7], k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右轮转 1 步: [7,1,2,3,4,5,6]
//向右轮转 2 步: [6,7,1,2,3,4,5]
//向右轮转 3 步: [5,6,7,1,2,3,4]
//
//
// 示例 2:
//
//
//输入：nums = [-1,-100,3,99], k = 2
//输出：[3,99,-1,-100]
//解释:
//向右轮转 1 步: [99,-1,-100,3]
//向右轮转 2 步: [3,99,-1,-100]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -2³¹ <= nums[i] <= 2³¹ - 1
// 0 <= k <= 10⁵
//
//
//
//
// 进阶：
//
//
// 尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
// 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
//
//
// Related Topics 数组 数学 双指针 👍 1887 👎 0

// 本题跟旋转数组相关
// 方法一: 使用额外空间，关键是放到新数组的 index := (i + k) % n
// 时间复杂度和空间复杂度都是 O(n)
// 注意 k 可能是比较大的, 可能大于数组的长度, 但取模后结果都一样
// 比如 len(nums) = 7, 那么 k = 3 和 k = 10 效果都是一样的
// 或者可以先把 k = k % 7
func rotate(nums []int, k int) {
	n := len(nums)
	arr := make([]int, n)
	for i := range nums {
		index := (i + k) % n
		arr[index] = nums[i]
	}

	copy(nums, arr)
}

// 方法二: 环状替换, 时间复杂度是O(n), 空间复杂度是O(1)
// 三个指针: start, cur 以及 next
// start: 指向环开始的位置
// cur: 指向正在处理的元素, 刚开始时cur = start
// next: 指向应该把cur位置的元素放到哪里
// 例子 nums = [1,2,3,4,5,6], k = 2
// 第一步: start = 0, cur 也等于0, 指向 i=0 的位置, next=(0+2) % 6 = 2 指向 i=2 的位置
//
//	所以i=0的元素放到i=2的位置, 但是不能直接覆盖了, 最好的办法就是通过一个中间变量prev替换
//	最后得到nums = [1,2,1,4,5,6], prev=3(保存原来i=2位置的元素), 然后让cur去到next的位置
//
// 第二步: start = 0, cur 等于 2, 指向 i=2 的位置, next=(2+2) % 6 = 4 指向 i=4 的位置
//
//	所以i=2的元素放到i=4的位置,i=2位置的元素已经存在prev里了,直接nums[1], prev = prev, nums[1] 即可
//	最后得到nums = [1,2,1,4,3,6], prev=5(保存原来i=4位置的元素), 然后让cur去到next的位置
//
// 第三步: start = 0, cur 等于4, 指向 i=4 的位置, next=(4+2) % 6 = 0 指向 i=0 的位置
//
//	所以i=4的元素放到i=0的位置,i=4位置的元素已经存在prev里了,直接nums[1], prev = prev, nums[1] 即可
//	最后得到nums = [5,2,1,4,3,6]
//
// 很明显: 从 i = 0 的位置开始, 最后又回到了 i = 0 的位置, 形成了一个换, 所以叫做环状替换
// 以此类推, 处理完一个环后让start++继续往下处理其它环即可
func rotate2(nums []int, k int) {
	n := len(nums)
	count := 0 // 用来存储处理了多少个元素
	start := 0
	for count < n {
		cur := start
		prev := nums[cur]
		// 进行环状替换
		for {
			next := (cur + k) % n
			nums[next], prev = prev, nums[next]
			cur = next
			count++

			// 回到了开始的位置, 表明处理完了一个环
			if cur == start {
				break
			}
		}

		start++
	}
}

// 方法三: 三次数组翻转 nums = [1,2,1,4,5,6,7]
// 第一次反转 [0, n-1]: [7,6,5,4,3,2,1]
// 第二次反转 [0, k-1]: [5,6,7,4,3,2,1]
// 第二次反转 [k, n-1]: [5,6,7,1,2,3,4]
func rotate3(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
