class Solution:
    def zigzagLevelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if not root:
            return []
        
        results = []
        queue = deque([root])
        level = 0
        while queue:
            length_level = len(queue)
            levels = []
            for _ in range(length_level):
                node = queue.popleft()
                levels.append(node.val)

                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
            if level % 2 != 0:
                left, right = 0, len(levels) - 1
                while left < right:
                    levels[left], levels[right] = levels[right], levels[left]
                    left += 1
                    right -= 1
            results.append(levels)
            level += 1
        return results
        