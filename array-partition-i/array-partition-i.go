import "sort"
func arrayPairSum(nums []int) int {
    sum := 0
    sort.Ints(nums)
    for i:=0;i<len(nums);i+=2 {
        if nums[i] < nums[i+1] {
            sum += nums[i]
        } else {
            sum += nums[i+1]
        }
    }
    return sum
}