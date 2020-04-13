package main;

import "testing"

func TestEqualNumberContiguousSubArray(t *testing.T) {
	tables := []struct {
		a []int
		exp int
	}{
		{ []int {}, 0},
		{ []int {1}, 0},
		{ []int {0,1}, 2},
		{ []int {0,1,1}, 2},
		{ []int {0,0,1}, 2},
		{ []int {0,0,1}, 2},
		{ []int {0,1,0,0,1}, 4},
		{ []int {0,1,1,0,1,1,1,0}, 4},
		{ []int {1,0,1,1,0,1,1,1,0}, 4},
    }
    
	for _, table := range tables {
		result:= findMaxLength(table.a);

        if result != table.exp {
			t.Errorf("%d, continuous subarray actual %d, should be %d", table.a, result, table.exp);
         }
	}
}