package threesum

import "testing"

func Test(t *testing.T) {
	t.Log(threeSum([]int{0, 0, 0, 0, 0}))
	t.Log(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	t.Log(threeSum([]int{0}))
}
