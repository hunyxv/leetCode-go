package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var res [][]int = [][]int{}

	sort.Ints(nums)

	for p1 := 0; p1<len(nums) - 3; p1++{
		if nums[p1] > 0 && target <= 0{
			return res
		}

		if p1 > 0 && nums[p1] == nums[p1 - 1] {
			continue
		}

		for p2 := p1 + 1; p2 < len(nums) - 2; p2++ {
			if p2 > p1 +1 && nums[p2] == nums[p2 - 1] {
				continue
			}

			p3, p4 := p2 + 1, len(nums) - 1

			for p3 < p4 {
				if nums[p1] + nums[p2] + nums[p3] == target - nums[p4] {
					res = append(res, []int{nums[p1], nums[p2], nums[p3], nums[p4]})
					
					for (p3 < p4 && nums[p3] == nums[p3+1]) {        // 去重
						p3++
					}
					for (p3 < p4 && nums[p4] == nums[p4-1]) {		// 去重
						p4--
					}
					p3++
					p4--
				}

				if p3 < p4 && nums[p1] + nums[p2] + nums[p3] < target - nums[p4] {
					p3++
				} else {
					p4--
				}
			}
		}
	}
	return res
}

func main() {
	nums := []int{1, 1, 1, 1}
	fmt.Println(fourSum(nums, 4))
}