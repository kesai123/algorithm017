package main


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var iRoot int
	for iRoot = range inorder {
		if inorder[iRoot] == preorder[0]{
			break
		}
	}
	return &TreeNode {
		preorder[0],
		buildTree(preorder[1:iRoot+1], inorder[0:iRoot]),
		buildTree(preorder[iRoot+1:], inorder[iRoot+1:]),
	}
}
