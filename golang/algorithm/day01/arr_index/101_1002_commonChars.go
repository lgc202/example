package arr_index

//给你一个字符串数组 words ，请你找出所有在 words 的每个字符串中都出现的共用字符（ 包括重复字符），并以数组形式返回。你可以按 任意顺序 返回答
//案。
//
//
//
// 示例 1：
//
//
//输入：words = ["bella","label","roller"]
//输出：["e","l","l"]
//
//
// 示例 2：
//
//
//输入：words = ["cool","lock","cook"]
//输出：["c","o"]
//
//
//
//
// 提示：
//
//
// 1 <= words.length <= 100
// 1 <= words[i].length <= 100
// words[i] 由小写英文字母组成
//
//
// Related Topics 数组 哈希表 字符串 👍 342 👎 0

// 方法: 将字符作为数组下标
func commonChars(words []string) []string {
	count := make([]int, 26)
	// 计算第一个字符串中各个字符的个数
	for _, str := range words[0] {
		// 'a'的 index 是 0，则 'h' 的 index 是 'h' - 'a'
		count[str-'a']++
	}

	for i := 1; i < len(words); i++ {
		tmpCount := make([]int, 26)
		for _, str := range words[i] {
			tmpCount[str-'a']++
		}

		// 在第一个字符出现的字符在第二个字符串中未必出现, 所以要取个数最小的
		for j := 0; j < 26; j++ {
			if count[j] > tmpCount[j] {
				count[j] = tmpCount[j]
			}
		}
	}

	// 最后遍历 count 数组生成结果
	var res []string
	for i := range count {
		for j := 0; j < count[i]; j++ {
			res = append(res, string(byte(i+'a')))
		}
	}

	return res
}
