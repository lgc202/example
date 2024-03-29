package day02

import "strconv"

//给定一个 无重复元素 的 有序 整数数组 nums 。
//
// 返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于
//nums 的数字 x 。
//
// 列表中的每个区间范围 [a,b] 应该按如下格式输出：
//
//
// "a->b" ，如果 a != b
// "a" ，如果 a == b
//
//
//
//
// 示例 1：
//
//
//输入：nums = [0,1,2,4,5,7]
//输出：["0->2","4->5","7"]
//解释：区间范围是：
//[0,2] --> "0->2"
//[4,5] --> "4->5"
//[7,7] --> "7"
//
//
// 示例 2：
//
//
//输入：nums = [0,2,3,4,6,8,9]
//输出：["0","2->4","6","8->9"]
//解释：区间范围是：
//[0,0] --> "0"
//[2,4] --> "2->4"
//[6,6] --> "6"
//[8,9] --> "8->9"
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 20
// -2³¹ <= nums[i] <= 2³¹ - 1
// nums 中的所有值都 互不相同
// nums 按升序排列
//
//
// Related Topics 数组 👍 268 👎 0

// 有序数组 + 双指针
func summaryRanges(nums []int) []string {
	// start 指向区间开头
	start := 0
	// end 指向区间结尾
	end := 0
	n := len(nums)

	var res []string
	for end < n {
		for end+1 < n && nums[end+1]-nums[end] == 1 {
			end++
		}

		if start == end {
			// 区间只有一个数字
			res = append(res, strconv.Itoa(nums[start]))
		} else {
			res = append(res, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[end]))
		}

		// 处理完一个区间后移动到下一个区间
		end++
		start = end
	}

	return res
}
