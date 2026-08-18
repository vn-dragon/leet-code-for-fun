class Solution:
    def goodNodes(self, root: TreeNode) -> int:
        self.res = 0
        def dfs(node, maxValue):
            if not node:
                return
            if node.val >= maxValue:
                self.res += 1
                maxValue = node.val
            dfs(node.left, maxValue)
            dfs(node.right, maxValue)
        dfs(root, root.val)
        return self.res