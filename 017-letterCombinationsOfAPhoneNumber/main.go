package main

import "fmt"

var keyboard map[byte][]string = map[byte][]string{
	'1': make([]string, 0),
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations (digits string) []string {
	if digits == "" {
		return []string{}
	}

	res := dfs(&digits, 0)
	return res
}

func dfs(digits *string, index int) []string {
	if index == len(*digits) - 1 {
		return keyboard[(*digits)[index]]
	}
	var res []string
	
	for _, c := range keyboard[(*digits)[index]] {
		for _, str := range dfs(digits, index + 1) {
			res = append(res, c + str)
		}
	}
	return res
}

func main(){
	res := letterCombinations("23")
	fmt.Println(res)	
}