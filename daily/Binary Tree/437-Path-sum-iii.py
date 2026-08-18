class Solution:
    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> int:
        if not root:
            return 0
        
        def dfs(node, currentSum):
            if not node:
                return 0
            
            currentSum += node.val
            count = 1 if currentSum == targetSum else 0
            left = dfs(node.left, currentSum)
            right = dfs(node.right, currentSum)
            return count + left + right
        return (dfs(root, 0) + self.pathSum(root.left, targetSum) + self.pathSum(root.right, targetSum))