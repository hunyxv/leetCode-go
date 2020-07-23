package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int = [][]int{}

	sort.Ints(nums)
	
	for p1 := 0; p1<len(nums); p1++ {
		if nums[p1] > 0 {
			return res
		}

		if p1 > 0 && nums[p1] == nums[p1-1] {
			continue
		}

		var p2, p3 int = p1 + 1, len(nums) - 1

		for p2 < p3 {
			if nums[p1] + nums[p2] == -nums[p3] {
				res = append(res, []int{nums[p1], nums[p2], nums[p3]})
			
				for (p2 < p3 && nums[p2] == nums[p2+1]) {        // 去重
					p2++
				}
				for (p2 < p3 && nums[p3] == nums[p3-1]) {		// 去重
					p3--
				}
				p2++
				p3--
			} else if nums[p1] + nums[p2] < -nums[p3] {			// 去重
				p2++
			} else {
				p3--
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1}))
}