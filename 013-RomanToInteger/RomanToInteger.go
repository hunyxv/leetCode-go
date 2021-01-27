package romantointeger

func romanToInt(s string) int {
    m := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
    l := len(s)
    var result int
    for i := 0; i < l; i++ {
        /*
        I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
        X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。 
        C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900
        */
        if i < l - 1 && m[s[i]] < m[s[i+1]] {
            result -= m[s[i]]
        } else {
            result += m[s[i]]
        }
    }
    return result
}