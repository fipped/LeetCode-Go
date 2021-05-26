package leetcode

func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    r, x, dx := make([][]rune, numRows), 0, 1
    for _, c := range s {
        r[x] = append(r[x], c)
        if dx == 1 {
            if x == numRows - 1 {
                dx = -1
            }
        } else if dx == -1 {
            if x == 0 {
                dx = 1
            }
        }
        x += dx
    }
    res := ""
    for _, row := range r {
        res += string(row)
    }
    return res
}
