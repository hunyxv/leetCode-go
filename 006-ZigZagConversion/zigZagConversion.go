package zigzagconversion

func convert(s string, numRows int) string {
	l := len(s)
	i := 0
	buf := make([][]byte, numRows)
	for i < l {
		for j := 0; j < numRows && i < l; j, i = j+1, i+1 {
			buf[j] = append(buf[j], s[i])
		}

		for j := 0; j < numRows-2 && i < l; j, i = j+1, i+1 {
			buf[numRows-j-2] = append(buf[numRows-j-2], s[i])
		}
	}

	var result string
	for i := range buf {
		result += string(buf[i])
	}
	return result
}
