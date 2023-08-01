package day02

//给你一个长度为 n 的整数数组
// nums ，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。
//
// 我们是这样定义一个非递减数列的： 对于数组中任意的 i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。
//
//
//
// 示例 1:
//
//
//输入: nums = [4,2,3]
//输出: true
//解释: 你可以通过把第一个 4 变成 1 来使得它成为一个非递减数列。
//
//
// 示例 2:
//
//
//输入: nums = [4,2,1]
//输出: false
//解释: 你不能在只改变一个元素的情况下将其变为非递减数列。
//
//
//
//
// 提示：
//
//
//
// n == nums.length
// 1 <= n <= 10⁴
// -10⁵ <= nums[i] <= 10⁵
//
//
// Related Topics 数组 👍 754 👎 0

// 本题跟有序数组相关
// 当当前元素比前一个元素小时, 说明递减了, 需要改变一个元素
// 改变元素时可以把 num[i] 增大一点让它跟num[i-1] 一样大, 也可以把 nums[i-1] 减小点让它跟num[i] 一样
// 情况一: 4,5,3, 当前i指向3, 而 i-1指向5, 由于nums[i] < nums[i-2], 不可能让num[i-1]减小成4,3,3,只能让num[i]增大成4,5,5
// 情况一: 4,3,5, 当前i指向3, 而 i-1指向5, 由于nums[i] > nums[i-2], 让nums[i-1]增大成4,5,5肯定可以保证是非递减序列
// 注意 i 是从 1 开始的
func checkPossibility(nums []int) bool {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			count++
			// 只能改一个元素
			if count > 1 {
				return false
			}

			// 可以改 第 i 个元素也可以改第 i-1 个元素
			// 具体看哪个 num[i] 和 nums[i-2]的大小关系
			if i-2 >= 0 && nums[i] < nums[i-2] {
				nums[i] = nums[i-1]
			} else {
				nums[i-1] = nums[i]
			}
		}
	}

	return true
}
