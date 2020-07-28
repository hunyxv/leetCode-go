package main

import (
	"fmt"
)

var pair map[rune]rune = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
}

type Stack []rune

func (s *Stack) push(c rune){
	*s = append(*s, c)
}

func (s *Stack) pop() (rune, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	c := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return c, true
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	var stack Stack
	for _, c := range s {
		if c == '{' || c == '(' || c == '[' {
			stack.push(c)
		} else if r, ok := stack.pop(); !ok || c != pair[r] {
			return false
		}
	}
	
	return len(stack) == 0
}

func main() {
	s := "(("
	fmt.Println(isValid(s))

	s = "[()]{}"
	fmt.Println(isValid(s))
}