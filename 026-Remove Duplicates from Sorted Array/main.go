package main

import "fmt"

// func removeDuplicates(nums []int) int {
// 	var tmp, count, ll int = 0, 1, len(nums)
// 	fmt.Println(nums)
// 	for i:= 1; i < ll; i++ {
// 		if nums[i] != nums[tmp] {
// 			ll = ll-(i-tmp-1)
// 			count++
// 			if i - tmp > 1 {
// 				nums = append(nums[:tmp+1], nums[i:]...)
// 			}
// 			tmp++
// 			i = tmp
// 		}
// 	}
// 	return count
// }

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	count, p, tmp := 1, 0, 0
	for i:=0; i<len(nums); i++ {
		if nums[i] != nums[tmp] {
			p++
			nums[p] = nums[i]
			tmp = i
			count++
		}
	}
	fmt.Println(nums[:count])
	return count
}

func main() {
	count := removeDuplicates([]int{1,1,2,3,3,4,5,5})
	fmt.Println(count)
}
