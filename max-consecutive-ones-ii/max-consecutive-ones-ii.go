func findMaxConsecutiveOnes(nums []int) int {
    max, count1, count2 := 0, 0, 0
    for _, num := range nums {
        if num == 0 {
            if count2 > max {
                max = count2
            }
            count2 = count1 + 1
            count1 = 0
        } else if num == 1 {
            count1++
            count2++
        } else {
            panic("non-binary number.")
        }
    }
    if count1 > max {
        max = count1
    }
    if count2 > max {
        max = count2
    }
    return max
}