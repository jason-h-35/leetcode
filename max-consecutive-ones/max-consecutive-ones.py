class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        total_max = 0
        current_max = 0
        for num in nums:
            if num == 1:
                current_max += 1
            else:
                total_max = max(total_max, current_max)
                current_max = 0
        return max(total_max, current_max)
                
        