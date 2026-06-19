class Solution:
    def deleteNode(self, root: Optional[TreeNode], key: int) -> Optional[TreeNode]:
        def balanceBST(leftNode, rightNode):
            if not rightNode:
                return leftNode
            rightNode.left = balanceBST(leftNode, rightNode.left)
            return rightNode

        if not root:
            return None

        if root.val == key:
            return balanceBST(root.left, root.right)

        if root.val > key:
            root.left = self.deleteNode(root.left, key)
        elif root.val < key:
            root.right = self.deleteNode(root.right, key)
        return root     