package main;

import "testing"

func TestMaximalSquare(t *testing.T) {
	tables := []struct {
		a [][]byte
		exp int
	}{
		{ [][]byte{}, 0},
		{ [][]byte{ {'1'}}, 1},
		{ [][]byte{ 
			{'0','1','1'},
			{'0','1','1'},
			{'0','1','0'},
        }, 4},        
        { [][]byte{ 
			{'0','1','0'},
			{'0','1','1'},
			{'0','1','1'},
        }, 4},
        { [][]byte{ 
			{'1','1','1'},
			{'1','1','1'},
			{'1','1','1'},
        }, 9},
        { [][]byte{ 
			{'1','1','1'},
			{'1','0','1'},
			{'1','1','1'},
        }, 1},
        { [][]byte {
        {'0','1','1','0','0','1','0','1','0','1'},
        {'0','0','1','0','1','0','1','0','1','0'},
        {'1','0','0','0','0','1','0','1','1','0'},
        {'0','1','1','1','1','1','1','0','1','0'},
        {'0','0','1','1','1','1','1','1','1','0'},
        {'1','1','0','1','0','1','1','1','1','0'},
        {'0','0','0','1','1','0','0','0','1','0'},
        {'1','1','0','1','1','0','0','1','1','1'},
        {'0','1','0','1','1','0','1','0','1','1'}}, 4},
    }
    
		    
	for _, table := range tables {

        duplicate := make([][]byte, len(table.a))
        for i := range table.a {
            duplicate[i] = make([]byte, len(table.a[i]))
            copy(duplicate[i], table.a[i])
        }

		result :=  maximalSquare(table.a);
		if result != table.exp {
            t.Errorf("%s, should have a square of area. Actual %d should be %d", duplicate, result, table.exp);
		 }
	}
}

