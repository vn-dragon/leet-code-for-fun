class Solution:
    def binaryTreePaths(self, root: Optional[TreeNode]) -> List[str]:
        results = []
        def dfs(node, path):
            if not node:
                return
            
            path.append(node.val)
            if not node.left and not node.right:
                leafPath = "->".join(str(val) for val in path)
                results.append(leafPath)
            dfs(node.left, path)
            dfs(node.right, path)

            path.pop()
        dfs(root, [])
        return results 