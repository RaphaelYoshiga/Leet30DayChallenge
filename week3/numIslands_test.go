package main

import "testing"

func TestNumIslands(t *testing.T) {
	tables := []struct {
		a [][]byte
		exp int
	}{
		{ [][]byte{}, 0},
		{ [][]byte{ {'1'}}, 1},
		{ [][]byte{ {'1','0','1'}}, 2},
		{ [][]byte{ {'1','1','1'}}, 1},
		{ [][]byte{ {'1','1','1','0','1'}}, 2},
		{ [][]byte{ 
			{'0','1','0'},
			{'0','1','1'},
			{'0','1','0'},
		}, 1},
		{ [][]byte{ 
			{'1','1','1','1','0'},
			{'1','1','0','1','0'},
			{'1','1','0','0','0'},
			{'0','0','0','0','0'},
		}, 1},
		{ [][]byte{ 
			{'0','1','0'},
			{'0','1','1'},
			{'0','1','0'},
		}, 1},
		{ [][]byte{ 
			{'1','1','1'},
			{'0','1','0'},
			{'1','1','1'},
		}, 1},
	}
		    
	for _, table := range tables {
		result :=  numIslands(table.a);
		if result != table.exp {
			t.Errorf("%s, invalid number of islands. actual %d should be %d", table.a, result, table.exp);
		 }
	}
}