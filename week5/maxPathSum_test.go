package main

import "testing"

func TestMaxPathSumBinaryTree(t *testing.T) {
	tables := []struct {
		x *TreeNode
		exp int
	}{
		{ &TreeNode{ Val: 1}, 1},
		{ &TreeNode{ Val: 1, Left: &TreeNode{ Val: 2}}, 3},
		{ &TreeNode{ Val: 1, Right: &TreeNode {Val: 3}, Left: &TreeNode{ Val: 2}}, 6},
		{ &TreeNode{ Val: 1, Right: &TreeNode {Val: 3}, Left: &TreeNode{ Val: -2}}, 4},
		{ &TreeNode{ Val: 9, 
			Left: &TreeNode{ Val: 6 },
			Right: &TreeNode{ Val: -3,
				Left: &TreeNode {Val: -6},
				Right: &TreeNode {Val: 2, 
						Left: &TreeNode {Val: 2, 
							Left: &TreeNode {Val: -5},
							Right: &TreeNode {Val: -5}}}}},16},

		// { &TreeNode{ Val: 1, Left: &TreeNode{ Val: 2, Left: &TreeNode { Val: 4, Left: &TreeNode { Val: 6}}}}, 3},
		// { &TreeNode{ Val: 1, 
		// 	Left: &TreeNode{ Val: 2, 
		// 		Left: &TreeNode { Val: 4 },
		// 		Right: &TreeNode { Val: 8, 
		// 			Left: &TreeNode { Val: 9}}},
		// 	Right: &TreeNode{ Val: 3 }},4},
	}
	for _, table := range tables {
        result := maxPathSum(table.x)
        if result != table.exp {
            t.Errorf("Wrong binary max path diameter! actual: %d, should be: %d", result, table.exp)
         }
	}
}
