import "math"
func thirdMax(nums []int) int {
    first, second, third := math.MinInt, math.MinInt, math.MinInt
    for _, num := range nums {
        if third < num && num < second {
            third = num
        } else if second < num && num < first {
            third = second
            second = num
        } else if first < num {
            third = second
            second = first
            first = num
        }
    }
    if third != math.MinInt {
        return third
    } else {
        return first
    }
}