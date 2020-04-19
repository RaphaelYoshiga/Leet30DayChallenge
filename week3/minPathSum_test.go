package main

import "testing"

func TestMinSumPath(t *testing.T) {
	tables := []struct {
		a [][]int
		exp int
	}{
		{ [][]int{}, 0},
		{ [][]int{ {1}}, 1},
		{ [][]int{ 
			{1,2,2},
			{1,2,2},
			{1,1,1},
		}, 5},
		{ [][]int{ 
			{0,0},
			{0,0},
		}, 0},
		{ [][]int{ 
			{1,1,1},
			{2,2,1},
			{2,2,1},
		}, 5},
		{ [][]int{ 
			{1,1,2},
			{2,1,2},
			{2,1,1},
		}, 5},
		{ [][]int{ 
			{1,0,0},
			{1,0,0},
			{1,1,0},
		}, 1},
		{ [][]int{ 
			{1,2,5},
			{3,2,1},
		}, 6},
	}
		    
	for _, table := range tables {
		result :=  minPathSum(table.a);
		if result != table.exp {
			t.Errorf("%d, invalid min path sum. actual %d should be %d", table.a, result, table.exp);
		 }
	}
}