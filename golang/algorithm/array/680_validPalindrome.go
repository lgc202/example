package array

func validPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	// 利用对撞双指针
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			// 删除左边或者右边的一个字符后仍然是回文串
			return check(s, left+1, right) || check(s, left, right-1)
		}
	}

	return true
}

func check(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}
