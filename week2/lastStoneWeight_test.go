package main

import "testing"

func TestLastStoneWeight(t *testing.T) {
	tables := []struct {
		a []int
		exp int
	}{
		{ []int {}, 0},
		{ []int {1}, 1},
		{ []int {3,2}, 1},
		{ []int {3,9}, 6},
		{ []int {1,9}, 8},
		{ []int {3,3,1}, 1},
		{ []int {1,3,3}, 1},
		{ []int {1,2,3,3}, 1},
		{ []int {3,7,2}, 2},
    }
    
	for _, table := range tables {
		result:= lastStoneWeight(table.a);

        if result != table.exp {
			t.Errorf("%d, last stone weight is wrong, actual %d, should be %d", table.a, result, table.exp);
         }
	}
}