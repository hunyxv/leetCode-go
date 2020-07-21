package main

import "fmt"

func longestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var left, right = 0, -1
	var result int 
	var freq [256]int
	
	for left < len(s) {
		if right + 1 < len(s) && freq[s[right+1]] == 0{
			freq[s[right + 1]]++
			right++
		}else{
			freq[s[left]]--
			left++
		}
		if right - left + 1 > result {
			result = right - left + 1
		}
	}
	return result
}

func main() {
	s := "rjegueawefdsk"
	fmt.Println(longestSubstring(s))
}