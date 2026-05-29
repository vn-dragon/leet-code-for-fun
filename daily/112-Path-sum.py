class Solution:
    def hasPathSum(self, root: Optional[TreeNode], targetSum: int) -> bool:
        if not root:
            return False
        def dfs(node, currentSum):
            if not node:
                return False
            currentSum += node.val
            if not node.left and not node.right:
                return currentSum == targetSum

            left = dfs(node.left, currentSum)
            right = dfs(node.right, currentSum)
             
            return left or right
        return dfs(root, 0)