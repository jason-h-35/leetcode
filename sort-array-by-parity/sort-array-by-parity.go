// O(n) time and O(1) space
// no using built in sort
// no creating a new array

func sortArrayByParity(nums []int) []int {
    i, j := 0, len(nums) - 1
    for i < j {
        for nums[i] % 2 == 0 && i < j {
            i++
        }
        for nums[j] % 2 == 1 && i < j {
            j--
        }
        nums[i], nums[j] = nums[j], nums[i]
    }
    return nums
}
