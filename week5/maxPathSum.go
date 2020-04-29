package main;

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
    if root == nil{
		return 0;
	}

	maxVal := root.Val;
	deepestNodeAt(root, &maxVal);
    return maxVal;
}

func deepestNodeAt(node *TreeNode, maxVal *int) int{
	if node == nil{
		return 0;
	}

	l :=max(deepestNodeAt(node.Left, maxVal), 0);
	r:= max(deepestNodeAt(node.Right, maxVal), 0);


	maxNodeSum := l + r + node.Val
	*maxVal = max(*maxVal, maxNodeSum);
	return max(l, r) + node.Val;
	
}

func max(a int, b int) int{
	if a > b{
		return a;
	}
	return b;
}

func main(){

	maxPathSum(&TreeNode{ Val: 9, 
		Left: &TreeNode{ Val: 6 },
		Right: &TreeNode{ Val: -3,
			Left: &TreeNode {Val: -6},
			Right: &TreeNode {Val: 2, 
					Left: &TreeNode {Val: 2, 
						Left: &TreeNode {Val: -5},
						Right: &TreeNode {Val: -5}}}}});
}