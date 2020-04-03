package main

import "testing"

func TestSumOfSubArray(t *testing.T) {
	tables := []struct {
		x []int
		exp int
	}{
		{ []int { 1,1 }, 2},
		{ []int { 1,3 }, 4},
		{ []int { 1,3,2 }, 6},
		{ []int { 1,-2, 7 }, 7},
		{ []int { 3,-2, 7 }, 8},
		{ []int { 1,3,2, -1 }, 6},
		{ []int { 1,3,2, -1, 9 }, 14},
		{ []int { 5,5,3,-10, 13 }, 16},
		{ []int { 5,5,3,-10, 13, -30, 90 }, 90},
		{ []int { 5,5,3,-10, 13, -30, 1 }, 16},
		{ []int { -1 }, -1},
		{ []int { -2, -1 }, -1},
    }
    
	for _, table := range tables {
        result := maxSubArray(table.x)
        if result != table.exp {
            t.Errorf("wrong sum! %d actual: %d, should be: %d", table.x, result, table.exp)
         }
	}
}
