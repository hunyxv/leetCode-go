package main

import "fmt"

func mostWater(heights []int) int {
	var v, tmp int
	var head, tail int = 0, len(heights) - 1

	for head < tail {
		if heights[head] < heights[tail] {
			tmp = (tail - head) * heights[head]
			head++
		} else {
			tmp = (tail - head) * heights[tail]
			tail--
		}
		if tmp > v {
			v = tmp
		}
	}
	return v
}

func main(){
	list := []int{1,11,6,2,5,4,8,8,7}
	fmt.Println(mostWater(list))
}