package double_pointer

import "strings"

//如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
//
// 字母和数字都属于字母数字字符。
//
// 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入: s = "A man, a plan, a canal: Panama"
//输出：true
//解释："amanaplanacanalpanama" 是回文串。
//
//
// 示例 2：
//
//
//输入：s = "race a car"
//输出：false
//解释："raceacar" 不是回文串。
//
//
// 示例 3：
//
//
//输入：s = " "
//输出：true
//解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
//由于空字符串正着反着读都一样，所以是回文串。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 2 * 10⁵
// s 仅由可打印的 ASCII 字符组成
//
//
// Related Topics 双指针 字符串 👍 656 👎 0

// 思路：使用对撞双指针
// 关键点:
// (1) 将字符串全部转换为小写或者大写
// (2) 判断字符是不是字母或者数字
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	left := 0
	right := len(s) - 1
	s = strings.ToLower(s)
	for left < right {
		if !isLetterOrDigit(s[left]) {
			left++
			continue
		}
		if !isLetterOrDigit(s[right]) {
			right--
			continue
		}

		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func isLetterOrDigit(b byte) bool {
	return ('a' <= b && b <= 'z') || ('0' <= b && b <= '9')
}
