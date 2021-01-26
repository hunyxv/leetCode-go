package longesubstring

func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }

    var result, left, right int
    var freq [256]int

    for right < len(s) {
        if freq[s[right]] == 0 {
            freq[s[right]]++
            right++
        } else {
            freq[s[left]]--
            left++
        }
        if right - left > result {
            result = right - left
        }
    }
    return result
}
