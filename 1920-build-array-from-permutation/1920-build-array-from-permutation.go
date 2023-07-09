func buildArray(nums []int) []int {
    ans := make([]int, len(nums))
    for i, x := range nums {
        ans[i] = nums[x]
    }
    return ans
}