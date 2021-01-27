package romantointeger

import "testing"

func Test(t *testing.T) {
	if romanToInt("IV") != 5 {
		t.Fail()
	}
}