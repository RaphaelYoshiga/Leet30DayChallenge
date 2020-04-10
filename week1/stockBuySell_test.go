package main

import "testing"

func TestBuySellStock(t *testing.T) {
	tables := []struct {
		x []int
		exp int
	}{
		{ []int { 2,1 }, 0},
		{ []int { 1,2 }, 1},
		{ []int { 2,4 }, 2},
		{ []int { 1,6 }, 5},
		{ []int { 1,6,9 }, 8},
		{ []int { 1,9,6 }, 8},
		{ []int { 2,1,9,6 }, 8},
		{ []int { 1,2,3,4,5 }, 4},
		{ []int { 7,6,4,3,1}, 0},
		{ []int { 7,1,5,3,6,4}, 7},		
		{ []int { 3,6,7,3,6,7}, 8},		
		{ []int { 2,1,2,0,1}, 2},		
		{ []int { }, 0},		
		
    }
    
	for _, table := range tables {
        result := maxProfit(table.x)
        if result != table.exp {
            t.Errorf("Wrong profit! %d actual: %d, should be: %d", table.x, result, table.exp)
         }
	}
}
