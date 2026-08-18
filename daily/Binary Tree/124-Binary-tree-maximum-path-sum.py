class Solution:
    def maxPathSum(self, root: Optional[TreeNode]) -> int:
        self.best = float('-inf')
        def dfs(node):
            if not node:
                return 0
            left = dfs(node.left)
            right = dfs(node.right)
            if left < 0:
                left = 0
            if right < 0:
                right = 0
            self.best = max(self.best, node.val + left + right)
            return node.val + max(left, right)
        dfs(root)
        return self.best