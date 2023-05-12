class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        ransomMap = {}
        magazineMap = {}
        for char in ransomNote:
            if char in ransomMap:
                ransomMap[char] += 1
            else:
                ransomMap[char] = 1
        for char in magazine:
            if char in magazineMap:
                magazineMap[char] += 1
            else:
                magazineMap[char] = 1
        for k in ransomMap:
            if k not in magazineMap:
                return False
            elif ransomMap[k] > magazineMap[k]:
                return False
        return True