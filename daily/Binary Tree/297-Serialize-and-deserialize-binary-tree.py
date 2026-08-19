class Codec:
    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        ans = []
        def dfs(node):
            if not node:
                ans.append("X")
                return
            ans.append(str(node.val))
            dfs(node.left)
            dfs(node.right)
        dfs(root)
        return ",".join(ans)
    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        preorder = iter(data.split(','))
        def build():
            val = next(preorder)
            if val is "X":
                return None
            node = TreeNode(int(val))
            node.left = build()
            node.right = build()
            return node
        return build()