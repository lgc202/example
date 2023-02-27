package dfs_bfs

// 给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按任意顺序返回。
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
// 示例 1：
// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
// 示例 2：
// 输入：digits = ""
// 输出：[]
//
// 示例 3：
// 输入：digits = "2"
// 输出：["a","b","c"]

func letterCombinationsDFS(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	table := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	var res []string

	// index 表示访问到第几个数字了，也可以认为是访问到树的第几层
	// path 表示从根节点到叶子节点的路径
	var dfs func(index int, path string)
	dfs = func(index int, path string) {
		// 到达叶子节点了， 就把这条选择的路径加到res中
		if len(path) == len(digits) {
			res = append(res, path)
			return
		}

		// 访问当前节点的所有子节点
		for _, val := range table[string(digits[index])] {
			dfs(index+1, path+string(val))
		}
	}

	dfs(0, "")
	return res
}

func letterCombinationsBFS(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	table := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	var queue []string
	// 根节点入队
	queue = append(queue, "")
	for i := range digits {
		number := string(digits[i])
		for _, val := range queue {
			// 出队
			queue = queue[1:]
			// 遍历子节点
			for _, c := range table[number] {
				queue = append(queue, val+string(c))
			}
		}
	}

	return queue
}
