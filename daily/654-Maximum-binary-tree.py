class Solution:
    def constructMaximumBinaryTree(self, nums: List[int]) -> Optional[TreeNode]:
        def dfs(arr: List[int]):
            if not arr:
                return None
            max_value = max(arr)
            index_maxVal = arr.index(max_value)
            node = TreeNode(max_value)
            left_prefix = arr[0:index_maxVal]
            right_suffix = arr[index_maxVal + 1: ]
            node.left = dfs(left_prefix)
            node.right = dfs(right_suffix)
            return node

        return dfs(nums)