package main;

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isValidSequence(root *TreeNode, arr []int) bool {

	return isValidPath(root, 0, arr);
}

func isValidPath(node *TreeNode, i int, arr[] int) bool{
	val:= arr[i];
	
	if node == nil || val != node.Val{
		return false;
	}

	if i == len(arr) -1{
		return node.Right == nil && node.Left == nil;
	}

	i++;
	return isValidPath(node.Left, i, arr) || isValidPath(node.Right, i, arr)
}