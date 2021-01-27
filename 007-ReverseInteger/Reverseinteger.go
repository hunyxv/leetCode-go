package reverseinteger

func reverse(x int) int {
	if x == 0 {
		return x
	}

	var result int
	for x != 0 {
		mid := x % 10
		x = x / 10
		result = result*10 + mid
	}
	if result > 1<<31-1 || result < -(1<<31) {
		return 0
	}

	return result
}
