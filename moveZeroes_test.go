package main

import "testing"

func TestMoveZeroes(t *testing.T) {
	tables := []struct {
		x []int;
		exp []int
	}{
		{[]int { 0,1 }, []int {1,0}},
		{[]int { 0,0 }, []int {0,0}},
		{[]int { 0,7 }, []int {7,0}},

		{[]int { 1,1,0 }, []int {1,1,0}},
		{[]int { 0,1,1 }, []int {1,1,0}},		
		{[]int { 0,0,1,1 }, []int {1,1,0,0}},
		{[]int { 0,0,0,1,1 }, []int {1,1,0,0,0}},
		{[]int { 1,0,2,0,3,0,4,0,5,0 }, []int {1,2,3,4,5,0,0,0,0,0}},
    }
    
	for _, table := range tables {
		moveZeroes(table.x)

		for i, num := range table.exp {
			if table.x[i] != num {
				t.Errorf("%d should have been %d", table.x, table.exp)
				break;
			 }
		}
	}
}