/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 //第一种方法： 暴力递归法
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root := preorder[0]
    tree := &TreeNode{
        Val : root,
    }
    var count int
    for i, v := range inorder {
        if preorder[0] == v {
            count = i 
            break
        }
    }
    tree.Left = buildTree(preorder[1:count+1], inorder[:count])
    tree.Right = buildTree(preorder[count+1:], inorder[count+1:])
    return tree
}