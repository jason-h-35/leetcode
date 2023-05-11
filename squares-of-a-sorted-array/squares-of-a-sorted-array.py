class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        output = []
        i, j = 0, len(nums) - 1
        while i <= j:
            if nums[i] ** 2 <= nums[j] ** 2:
                output.append(nums[j] ** 2)
                j -= 1
            else:
                output.append(nums[i] ** 2)
                i += 1
        return reversed(output)