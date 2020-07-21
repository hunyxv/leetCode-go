package main

import "fmt"


func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int)

	for index, value := range nums {
		dif := target - value

		if dIndex, ok := tmp[dif]; ok {
			return []int{dIndex, index}
		}
		tmp[value] = index
	}
	return nil
}

func main() {
	res := twoSum([]int{2,7,11,15}, 18)
	fmt.Println(res)
}