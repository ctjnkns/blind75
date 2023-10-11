/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return valid(root, nil, nil)
}

func valid(node, left, right *TreeNode) bool {
    if node == nil {
        return true //base case
    }
    if (right != nil && node.Val >= right.Val) || (left != nil && node.Val <= left.Val) {
        return false
    }
    return valid(node.Left, left, node) && valid(node.Right, node, right)
}
