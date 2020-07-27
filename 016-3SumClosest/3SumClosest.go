package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	var res int = 1 <<63-1

	sort.Ints(nums)
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	for p1 := 0; p1 < len(nums); p1++ {
		if p1 > 0 && nums[p1] == nums[p1-1] {
			continue
		}
		
		for p2, p3 := p1 + 1, len(nums) - 1; p2 < p3; {
			sum := nums[p1] + nums[p2] + nums[p3]
			
			if sum < target {
				p2++
			} else if sum > target {
				p3--
			} else {
				return sum
			}
			if abs(sum - target) < abs(res - target) {
				res = sum
			}
		}
	}
	return res
}


func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}