import "sort"
func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return 1
    }
    sort.Ints(nums)
    seqlen:= 1
    maxseqlen:= 1
    for i:= 1; i < len(nums); i++ {
        if nums[i-1] + 1 == nums[i] {
            seqlen++
            if seqlen > maxseqlen {
                maxseqlen = seqlen
            }
        } else if nums[i-1] != nums[i] {
            seqlen = 1
        }
    }
    return maxseqlen
}