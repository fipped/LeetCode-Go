package leetcode

func myAtoi(s string) int {
    i, l := 0, len(s)
    for ; i < l; i++ {
        if s[i] != ' ' {
            break
        }
    }
    
    d := 1
    if i < l {
        if s[i] == '-' {
            d = -1
            i++
        } else if s[i] == '+' {
            i++
        }
    }

    num := 0
    for ; i < l; i++ {
        if !isDigit(s[i]) {
            break
        }
        num = num * 10 + int(s[i] - '0')
        if num*d > (1 << 31) - 1 {
            return (1 << 31) - 1
        }
        if num*d < -(1 << 31) {
            return -(1 << 31)
        }
    }
    return num*d
}

func isDigit(c byte) bool {
    return c - '0' >= 0 && c - '0' <= 9
}
