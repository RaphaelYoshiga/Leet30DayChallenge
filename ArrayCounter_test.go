package main

import "testing"

func TestCountPlusOneInArray(t *testing.T) {
	tables := []struct {
		x []int
		exp int
	}{
		{ []int { 1,1,3,3,5,5,7 }, 0},
		{ []int { 1,2 }, 1},
		{ []int { 1,1,2,2}, 2},
		{ []int { 1,3,2,3,5,0 }, 3},
    }
    
	for _, table := range tables {
        result := countElements(table.x)
        if result != table.exp {
            t.Errorf("wrong sum! %d actual: %d, should be: %d", table.x, result, table.exp)
         }
	}
}
