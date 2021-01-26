package longestpalindromicsubstring


//  中心扩展算法
func longestPalindrome(s string) string {
	var start, end int
	for i := 0; i < len(s); i++ {
		left1, right1 := expandToBothSides(s, i, i)
		left2, right2 := expandToBothSides(s, i, i+1)

		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandToBothSides(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}


// 使用动态规划
func longestPalindrome2(s string) string {
    size := len(s)
    pb := make([][]bool, size)
    for i := range pb {
        pb[i] = make([]bool, size)
    }

    result := ""
    for l := 0; l < size; l++ {
        for i := 0; l + i < size; i++ {
            j := i + l
            if j == i {
                pb[i][j] = true
            } else if j - i == 1 && s[i] == s[j] {
                pb[i][j] = true
            } else {
                if s[i] == s[j] {
                    pb[i][j] = pb[i+1][j-1]
                }
            }
            if pb[i][j] && len(result) < j - i + 1 {
                result = s[i: j + 1]
            }
        }
    }
    return result
}
