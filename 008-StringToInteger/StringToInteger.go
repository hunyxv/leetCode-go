package stringtointeger

func myAtoi(s string) int {
	var result int
	var flag bool
	var i int

	for ; i < len(s); i++ {
		if s[i] == ' ' { // ascii 32 表示空格符
			continue
		}
		break
	}

	for ; i < len(s); i++ {
		if s[i] == ' ' { // ascii 32 表示空格符
			return 0
		}
		if s[i] == '-' {
			flag = true
			if i > 0 && s[i-1] == '+' {
				return 0
			}
			i++
			break
		}
		if s[i] == '+' {
			if flag {
				return 0
			}
			i++
			break
		}
		break
	}

	for ; i < len(s); i++ {
		if s[i] == ' ' || s[i] > '9' || s[i] < '0' {
			break
		}

		if flag {
			result = result*10 + (int(s[i]-'0') * -1)
			if result < -(1 << 31) {
				return -(1 << 31)
			}
		} else {
			result = result*10 + int(s[i]-'0')
			if result > 1<<31-1 {
				return 1<<31 - 1
			}
		}
	}

	return result
}
