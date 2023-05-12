# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from collections import deque
class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        levels = []
        # trivial case. root empty or None
        if not root:
            return levels
        level = 0
        # you HAVE to populate the queue with something for your loop to start with
        # it's the base case...
        queue = deque([root])
        # checking for empty is so easy... wtf.
        while queue: # until queue is empty...
            levels.append([]) # append empty list?
            for _ in range(len(queue)): # since we're adding shit to the queue
                node = queue.popleft()
                levels[level].append(node.val)
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
            level += 1 # outer, kinda...
        return levels