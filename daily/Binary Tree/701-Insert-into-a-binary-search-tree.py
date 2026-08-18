class Solution:
    def insertIntoBST(self, root: Optional[TreeNode], val: int) -> Optional[TreeNode]:
        def dfs(node, value):
            if not node:
                newNode = TreeNode(value)
                return newNode

            if node.val < value:
                node.right = dfs(node.right, value)
            elif node.val > value:
                node.left = dfs(node.left, value)
            return node
        return dfs(root, val)