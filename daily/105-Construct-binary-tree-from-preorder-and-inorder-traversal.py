class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        if not preorder:
            return None

        rootVal = preorder[0]
        node = TreeNode(rootVal)

        mid = inorder.index(rootVal)
        left_inorder = inorder[:mid]
        right_inorder = inorder[mid + 1: ]

        left_preorder_size = len(left_inorder)
        left_preorder = preorder[1:1+left_preorder_size]
        right_preorder = preorder[left_preorder_size + 1: ]

        node.left = self.buildTree(left_preorder, left_inorder)
        node.right = self.buildTree(right_preorder, right_inorder)

        return node