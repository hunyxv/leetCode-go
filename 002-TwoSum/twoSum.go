package twoSum

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

func twoSum2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		sub := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if sub == nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}

// func main() {}
