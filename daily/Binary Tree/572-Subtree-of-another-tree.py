class Solution:
    def isSubtree(self, root: Optional[TreeNode], subRoot: Optional[TreeNode]) -> bool:
        def dfs(root, subRoot):
            if not root or not subRoot:
                return False
            if isSameTree(root, subRoot):
                return True
            return dfs(root.left, subRoot) or dfs(root.right, subRoot)

        def isSameTree(node1, node2):
            if not node1 and not node2:
                return True
            if not node1 or not node2:
                return False
            return (node1.val == node2.val and isSameTree(node1.left, node2.left) and isSameTree(node1.right, node2.right))
        
        return dfs(root, subRoot)