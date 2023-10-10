/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if subRoot == nil {
        return true //edge case
    }
    if root == nil {
        return false //edge case
    }
    if sameTree(root, subRoot) {
        return true //base case
    }
    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func sameTree(s *TreeNode, t *TreeNode) bool {
    if s == nil && t == nil {
        return true
    }
    if s == nil || t == nil || s.Val != t.Val {
        return false
    }
    return sameTree(s.Left, t.Left) && sameTree(s.Right, t.Right)
}
