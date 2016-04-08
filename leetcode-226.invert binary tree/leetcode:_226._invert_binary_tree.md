#leetcode: 226. Invert Binary Tree

Invert a binary tree.

```

     4
   /   \
  2     7
 / \   / \
1   3 6   9
```
to

```
     4
   /   \
  7     2
 / \   / \
9   6 3   1
```

没啥好说的

```
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    
    root.Left, root.Right = root.Right, root.Left
    
    if root.Left != nil {
        invertTree(root.Left)
    }
    
    if root.Right != nil {
        invertTree(root.Right)
    }
    
    return root
}
```

那就黑一把谷歌吧233

> Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so fuck off.

虽然是事实，但权当一乐