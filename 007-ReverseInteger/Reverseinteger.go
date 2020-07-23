package main

import "fmt"

func reverseInteger(num int) int {
	var res int
	for num != 0 {
		res = res * 10 + num % 10
		num /= 10
	}
	if res < -(1<<31) || res > 1<<31-1 {
		return 0
	}
	return res
}

func main() {
	fmt.Println(reverseInteger(4634397))
	fmt.Println(reverseInteger(4637194397))
}