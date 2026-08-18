class Solution:
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        results = []
        def dfs(node: Optional[TreeNode]):
            if not node:
                return
            dfs(node.left)
            results.append(node.val)
            dfs(node.right)
        dfs(root)
        return results[k - 1]