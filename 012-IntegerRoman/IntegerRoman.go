package integerroman 

func intToRoman(num int) string {
    value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

    var result string
    for i:=0; i<13; i++ {
        for num >= value[i] {
            num -= value[i]
            result += roman[i]
        }
	}
	return result
}

func intToRoman2(num int) string {
	thousands := []string{"", "M", "MM", "MMM"}
    hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
    tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
    ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
    return thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
}