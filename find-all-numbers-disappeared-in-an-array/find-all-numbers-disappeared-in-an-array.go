func findDisappearedNumbers(nums []int) []int {
    counts := make([]int, len(nums)+1)
    for _, num := range nums {
        counts[num]++
    }
    missing := make([]int, len(nums))
    n := 0
    for num, count := range counts {
        if count == 0 && num != 0{
            missing[n] = num
            n++
        }
    }
    return missing[:n]
    
}