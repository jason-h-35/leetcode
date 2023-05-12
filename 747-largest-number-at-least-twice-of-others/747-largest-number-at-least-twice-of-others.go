func dominantIndex(nums []int) int {
    max := 0
    imax := -1
    for i,num := range nums {
        if num > max {
            max = num
            imax = i
        }
    }
    for i,num := range nums {
        if 2 * num > max && imax != i {
            return -1
        }
    }
    return imax
}