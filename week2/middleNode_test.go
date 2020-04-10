package main

import "testing"

func TestFetchMiddleNode(t *testing.T) {
	tables := []struct {
		x *ListNode
		exp int
	}{
		{ &ListNode{ Val: 1}, 1},
		{ &ListNode{ Val: 1, Next: &ListNode{ Val: 2}}, 2},
		{ &ListNode{ Val: 1, Next: &ListNode{ Val: 2, Next: &ListNode{ Val: 3}}}, 2},
		
    }
    
	for _, table := range tables {
        result := middleNode(table.x)
        if result.Val != table.exp {
            t.Errorf("Wrong node! actual: %d, should be: %d", result.Val, table.exp)
         }
	}
}
