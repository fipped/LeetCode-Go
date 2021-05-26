package leetcode

// 解法一 Manacher's algorithm，时间复杂度 O(n)，空间复杂度 O(n)
func longestPalindrome(s string) string {
    // RL[i]:    以预处理字符串下标 i 为中心的回文半径(奇数长度时不包括中心)
    // p:        通过中心扩散的方式能够扩散的最右边的下标
    // maxLen:   记录最长回文串的半径
    // start:    记录最长回文串在起始串 s 中的起始下标
    t := make([]rune, len(s) * 2 + 1)
    RL := make([]int, len(s) * 2 + 1)
    t[0] = '#'
    for i, c := range s {
        t[2 * i + 1] = c
        t[2 * i + 2] = '#'
    }
    p, start, maxLen := 0, 0, 0
    for i, _ := range t {
        if RL[p] + p > i {
            RL[i] = min(RL[2 * p - i], RL[p] + p - i)
        }
        // 中心扩散法求 RL[i]
        for ; i + RL[i] + 1 < len(t) && i - RL[i] - 1 >= 0 && t[i + RL[i] + 1] == t[i - RL[i] - 1]; RL[i]++ {
        }
        // 更新 p，它是遍历过的 i 的 i + RL[i] 的最大者
        if i + RL[i] > p + RL[p] {
            p = i
        }
        // 记录最长回文子串的长度和相应它在原始字符串中的起点
        if RL[i] > maxLen {
            maxLen = RL[i]
            // 偶数长度：#b#a#a# (i=4 即第3个#为中心，maxLen=2，起点为 1）
            // 奇数长度：#b#a#a#a#（i=5 即第2个a为中心，maxLen=3，起点为 1）
            start = (i - maxLen) / 2 
        }
    }
    return s[start:start+maxLen]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

// 解法二 滑动窗口，时间复杂度 O(n^2)，空间复杂度 O(1)
func longestPalindrome1(s string) string {
    if len(s) == 0 {
        return ""
    }
    left, right, pl, pr := 0, -1, 0, 0
    for left < len(s) {
        // 移动到相同字母的最右边（如果有相同字母）
        for right+1 < len(s) && s[left] == s[right+1] {
            right++
        }
        // 找到回文的边界
        for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
            left--
            right++
        }
        if right-left > pr-pl {
            pl, pr = left, right
        }
        // 重置到下一次寻找回文的中心
        left = (left+right)/2 + 1
        right = left
    }
    return s[pl : pr+1]
}

// 解法三 中心扩散法，时间复杂度 O(n^2)，空间复杂度 O(1)
func longestPalindrome2(s string) string {
    res := ""
    for i := 0; i < len(s); i++ {
        res = maxPalindrome(s, i, i, res)
        res = maxPalindrome(s, i, i+1, res)
    }
    return res
}

func maxPalindrome(s string, i, j int, res string) string {
    sub := ""
    for i >= 0 && j < len(s) && s[i] == s[j] {
        sub = s[i : j+1]
        i--
        j++
    }
    if len(res) < len(sub) {
        return sub
    }
    return res
}

// 解法四 DP，时间复杂度 O(n^2)，空间复杂度 O(n^2)
func longestPalindrome3(s string) string {
    res, dp := "", make([][]bool, len(s))
    for i := 0; i < len(s); i++ {
        dp[i] = make([]bool, len(s))
    }
    for i := len(s) - 1; i >= 0; i-- {
        for j := i; j < len(s); j++ {
            dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
            if dp[i][j] && (res == "" || j-i+1 > len(res)) {
                res = s[i : j+1]
            }
        }
    }
    return res
}
