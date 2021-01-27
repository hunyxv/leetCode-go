package threesum

import "sort"

func threeSum(nums []int) [][]int {
    result := make([][]int, 0)
    sort.Ints(nums)
    var ones, second, third  int
    for ones = 0 ; ones < len(nums)-2; ones++ {
        if nums[ones] > 1 {
            return result
        }

        if ones > 0 && nums[ones] == nums[ones-1] {
            continue
        }

        second, third = ones + 1, len(nums) - 1

        for second < third {
            if nums[ones] + nums[second] == -nums[third] {
                result = append(result, []int{nums[ones], nums[second], nums[third]})
                for second < third && nums[second] == nums[second+1] {
                    second++
                }
                for third > second && nums[third] == nums[third-1] {
                    third--
                }
                second++
                third--
            } else if nums[ones] + nums[second] < -nums[third] {
                second++
            } else {
                third--
            }
        }
    }
    return result
}