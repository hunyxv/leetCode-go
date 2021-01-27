package containerwithmostwater

func maxArea(height []int) int {
    var left, right int = 0, len(height)-1
    var tmp, cap int
    for left < right {
        if height[left] < height[right] {
            tmp = height[left] * (right - left)
            left++
        } else {
            tmp = height[right] * (right - left)
            right--
        }
        if tmp > cap {
            cap = tmp
        }
    }
    return cap
}