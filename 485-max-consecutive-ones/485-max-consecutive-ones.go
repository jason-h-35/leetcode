func findMaxConsecutiveOnes(nums []int) int {
    max := 0
    curmax := 0
    for _, num := range nums {
        if num == 1 {
            curmax++
            if curmax > max {
                max = curmax
            }
        } else {
            curmax = 0
        }
    }
    return max
}