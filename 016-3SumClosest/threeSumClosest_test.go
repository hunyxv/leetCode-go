package threesumclosest

import "testing"

func Test(t *testing.T) {
	if threeSumClosest([]int{-3,-2,-5,3,-4}, -1) != -2 {
		t.Fail()
	}
}