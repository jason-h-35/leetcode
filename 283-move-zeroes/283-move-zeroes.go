func moveZeroes(nums []int)  {
    n := 0
    for i,num := range nums {
        if num != 0 {
            nums[n] = nums[i]
            n++
        }
    }
    for i:=n;i<len(nums);i++ {
        nums[i] = 0
    }
}