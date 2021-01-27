package threesumclosest

import "sort"

func threeSumClosest(nums []int, target int) int {
	var result = 1<<32 - 1
	var ones, second, third, sum int

	sort.Ints(nums)
	for ones = 0; ones < len(nums)-2; ones++ {
		if ones > 0 && nums[ones] == nums[ones-1] {
			continue
		}

		second, third = ones+1, len(nums)-1
		for second < third {
			sum = nums[ones] + nums[second] + nums[third]
			if abs(sum-target) < abs(result-target) {
				result = sum
			}

			if sum < target {
				second++
			} else if sum > target {
				third--
			} else {
				return sum
			}
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
