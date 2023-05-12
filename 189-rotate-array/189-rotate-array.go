func rotate(nums []int, k int)  {
    k = k % len(nums) // avoid redundant shifting
    cut := len(nums)-k // define where to cut, k from the end of the array
    nums2 := make([]int, len(nums)); copy(nums2, nums) // copy to num2
    nums2 = append(nums2[cut:], nums2[:cut]...)
    for i, _ := range nums2 {
        nums[i] = nums2[i]
    }
}