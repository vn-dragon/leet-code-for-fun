package daily

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil || subRoot == nil {
        return false
    }
    if isSameTree(root, subRoot) {
        return true
    }
    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(node1 *TreeNode, node2 *TreeNode) bool {
    if node1 == nil && node2 == nil {
        return true
    }
    if node1 == nil || node2 == nil {
        return false
    }
    if node1.Val != node2.Val {
        return false
    }
    left := isSameTree(node1.Left, node2.Left)
    right := isSameTree(node1.Right, node2.Right)
    return left && right
}