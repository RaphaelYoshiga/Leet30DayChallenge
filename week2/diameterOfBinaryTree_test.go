package main

import "testing"

func TestBinaryTreeDiameter(t *testing.T) {
	tables := []struct {
		x *TreeNode
		exp int
	}{
		{ &TreeNode{ Val: 1}, 0},
		{ &TreeNode{ Val: 1, Left: &TreeNode{ Val: 2}}, 1},
		{ &TreeNode{ Val: 1, Right: &TreeNode {Val: 3}, Left: &TreeNode{ Val: 2}}, 2},
		{ &TreeNode{ Val: 1, Left: &TreeNode{ Val: 2, Left: &TreeNode { Val: 4, Left: &TreeNode { Val: 6}}}}, 3},
		{ &TreeNode{ Val: 1, 
			Left: &TreeNode{ Val: 2, 
				Left: &TreeNode { Val: 4 },
				Right: &TreeNode { Val: 8, 
					Left: &TreeNode { Val: 9}}},
			Right: &TreeNode{ Val: 3 }},4},
	}
	for _, table := range tables {
        result := diameterOfBinaryTree(table.x)
        if result != table.exp {
            t.Errorf("Wrong binary tree diameter! actual: %d, should be: %d", result, table.exp)
         }
	}
}
