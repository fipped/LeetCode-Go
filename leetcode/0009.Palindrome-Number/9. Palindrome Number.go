package leetcode

import "strconv"

// 解法一
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    s := make([]int, 0, 9)
    for x > 0 {
        s = append(s, x % 10)
        x /= 10
    }
    l := len(s)
    for i := 0; i < l/2; i++ {
        if s[i] != s[l-i-1] {
            return false
        }
    }
    return true
}

// 解法二 数字转字符串
func isPalindrome2(x int) bool {
    if x < 0 {
        return false
    }
    s := strconv.Itoa(x)
    l := len(s)
    for i := 0; i < l/2; i++ {
        if s[i] != s[l-i-1] {
            return false
        }
    }
    return true
}
