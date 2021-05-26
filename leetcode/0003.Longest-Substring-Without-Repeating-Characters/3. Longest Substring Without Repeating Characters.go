package leetcode

// 解法一 枚举每个位置为终点的最长无重复字母的子串长度，同时维护每个字母最后出现的位置，可更新起点为 min(起点, 上次当前字母出现位置+1）
func lengthOfLongestSubstring2(s string) int {
    letterPos := make(map[rune]int)
    maxLen := 0
    start := 0
    for i, letter := range s {
        if oldPos, ok := letterPos[letter]; ok {
            if oldPos + 1 > start {
                start = oldPos + 1
            }
        }
        length := i - start + 1
        if length > maxLen {
            maxLen = length
        }
        letterPos[letter] = i
    }
    return maxLen
}

// 解法二 滑动窗口
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var exist [256]bool
	result, left, right := 0, 0, 0
	for left < len(s) && right < len(s) {
		if !exist[s[right]] {
			exist[s[right]] = true
			right++
		} else {
			exist[s[left]] = false
			left++
		}
		result = max(result, right-left)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
