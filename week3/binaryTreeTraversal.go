package main;

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func bstFromPreorder(preorder []int) *TreeNode {

	if len(preorder) == 0{
		return nil;
	}
	
	root:= &TreeNode{  Val: preorder[0]};

	fo(1, preorder, root, -999, 99999);
    return root;
}


func fo(i int, preorder []int, node *TreeNode, min int, max int) int{
	if i >= len(preorder) || preorder[i]<min || preorder[i]>max{
		return i;
	}

	n := preorder[i];
	newNode := &TreeNode{ Val: n};
	if n < node.Val{
		node.Left = newNode;

		i++;
		i = fo(i, preorder, newNode, min, node.Val -1);
	}
	
	if i >= len(preorder) || preorder[i]<min || preorder[i]>max{
		return i;
	}

	n = preorder[i];
	newNode = &TreeNode{ Val: n};
	node.Right = newNode;
	i++;
	return fo(i, preorder, newNode, node.Val+1, max)
}

func main(){
	bstFromPreorder([]int {8,5,1,7,10,12});
}