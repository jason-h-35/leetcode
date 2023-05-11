func pivotIndex(nums []int) int {
    rightsum := sum(nums)
    leftsum := 0
    prevpivot := 0
    for i, pivot := range nums {
        leftsum += prevpivot
        rightsum -= pivot
        prevpivot = pivot
        if leftsum == rightsum {
            return i
        }
    }
    return -1
}

func sum(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return sum
}