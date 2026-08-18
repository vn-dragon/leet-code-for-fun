class Solution:
    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        def dfs(node):
            if not node:
                return 0, True
            leftHeight, leftBalance = dfs(node.left)
            rightHeight, rightBalance = dfs(node.right)

            currentBalance = leftBalance and rightBalance and abs(leftHeight - rightHeight) <= 1
            currentHeight = 1 + max(leftHeight, rightHeight)
            return (currentHeight, currentBalance)
        return dfs(root)[1]