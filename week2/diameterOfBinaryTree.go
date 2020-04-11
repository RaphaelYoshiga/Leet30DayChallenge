
package main;

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil{
		return 0;
	}

	diameter = 0;
	deepestNodeAt(root);
    return diameter;
}

func max( a int, b int) int{
	if a > b{
		return a;
	}
	return b;
}

var diameter int

func deepestNodeAt(node *TreeNode) int{
	if node == nil{
		return 0;
	}

	l :=deepestNodeAt(node.Left);
	r:= deepestNodeAt(node.Right);

	diameter = max(diameter, l+r);
	return max(l, r) + 1;
	
}