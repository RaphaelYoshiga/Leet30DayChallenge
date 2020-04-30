package main

import "testing"

func TestValidSequenceInBinaryTree(t *testing.T) {
	tables := []struct {
		x *TreeNode
		b []int
	}{
		{ &TreeNode{ Val: 1}, []int { 1 }},
		{ &TreeNode{ Val: 1, Left: &TreeNode{ Val: 2}}, []int { 1, 2 }},
		// { &TreeNode{ Val: 1, Left: &TreeNode{ Val: 2, Left: &TreeNode { Val: 4, Left: &TreeNode { Val: 6}}}}, 3},
		// { &TreeNode{ Val: 1, 
		// 	Left: &TreeNode{ Val: 2, 
		// 		Left: &TreeNode { Val: 4 },
		// 		Right: &TreeNode { Val: 8, 
		// 			Left: &TreeNode { Val: 9}}},
		// 	Right: &TreeNode{ Val: 3 }},4},
	}
	for _, table := range tables {
        result := isValidSequence(table.x, table.b)
        if !result {
            t.Errorf("%d is a valid path in three but it was not", table.b)
         }
	}
}

func TestInValidSequenceInBinaryTree(t *testing.T) {
	tables := []struct {
		x *TreeNode
		b []int
	}{
		{ &TreeNode{ Val: 1}, []int { 3 }},
		{ &TreeNode{ Val: 1, Left: &TreeNode{ Val: 2}}, []int { 1, 3 }},
		// { &TreeNode{ Val: 1, Left: &TreeNode{ Val: 2, Left: &TreeNode { Val: 4, Left: &TreeNode { Val: 6}}}}, 3},
		// { &TreeNode{ Val: 1, 
		// 	Left: &TreeNode{ Val: 2, 
		// 		Left: &TreeNode { Val: 4 },
		// 		Right: &TreeNode { Val: 8, 
		// 			Left: &TreeNode { Val: 9}}},
		// 	Right: &TreeNode{ Val: 3 }},4},
	}
	for _, table := range tables {
        result := isValidSequence(table.x, table.b)
        if result {
            t.Errorf("%d is not a valid path but should not be", table.b)
         }
	}
}
