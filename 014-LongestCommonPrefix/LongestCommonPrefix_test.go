package longestcommonprefix

import "testing"

func Test(t *testing.T) {
	if longestCommonPrefix([]string{"flower","flow","flight"}) != "fl" {
		t.Fail()
	}

	if longestCommonPrefix([]string{"dog","racecar","car"}) != "" {
		t.Fail()
	}
}