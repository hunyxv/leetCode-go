package palindromenumber

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    tmp := x
    var v int
    for tmp !=0 {
        v = v * 10 + tmp % 10
        tmp = tmp/10
    }
    return x == v
}