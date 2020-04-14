
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

	diameter := 0;
	deepestNodeAt(root, &diameter);
    return diameter;
}

func deepestNodeAt(node *TreeNode, diameter *int) int{
	if node == nil{
		return 0;
	}

	l :=deepestNodeAt(node.Left, diameter);
	r:= deepestNodeAt(node.Right, diameter);

	*diameter = max(*diameter, l+r);
	return max(l, r) + 1;
	
}