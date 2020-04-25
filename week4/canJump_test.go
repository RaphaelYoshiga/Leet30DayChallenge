package main;

import "testing"

func TestJumps(t *testing.T) {
	tables := []struct {
		a []int
		exp bool
	}{
		{ []int {1,1}, true},
		{ []int {2,0}, true},
		{ []int {2,0,1}, true},
		{ []int {1,1,1}, true},
		{ []int {2,5,0,0}, true},
		{ []int {1,1,1,1}, true},
		{ []int {2,3,1,1,4}, true},
		{ []int {2,0,1,1}, true},
		{ []int {0,1}, false},
		{ []int {3,2,1,0,4}, false},
    }
    
	for _, table := range tables {
		result:= canJump(table.a);

        if result != table.exp {
			t.Errorf("%d, %t should be %t", table.a, result, table.exp);
         }
	}
}

