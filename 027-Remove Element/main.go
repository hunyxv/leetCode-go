package main

import "fmt"

func removeElement(nums []int, val int) int {
	var p, count int = -1, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			count++
			p++
			nums[p] = nums[i]
		}
	}
	fmt.Println(nums[:count])
	return count
}

func main() {
	count := removeElement([]int{2,0,1,2,2,3,0,4,2}, 2)
	fmt.Println(count)
}