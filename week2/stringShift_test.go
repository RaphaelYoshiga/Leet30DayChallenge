
package main
import "testing"

func TestStringShifts(t *testing.T) {
	tables := []struct {
		a string
		b [][]int
		exp string
	}{
		{ "abcd", [][]int {{1,1}}, "dabc"},
		{ "abcd", [][]int {{1,2}}, "cdab"},
		{ "abcd", [][]int {{1,3}}, "bcda"},
		{ "abcd", [][]int {{1,5}}, "dabc"},
		{ "abcd", [][]int {{0,1}}, "bcda"},
		{ "abcd", [][]int {{0,3}}, "dabc"},
		{ "abcd", [][]int {{0,9}}, "bcda"},
		{ "abcd", [][]int {{0,3}, {1,3}}, "abcd"},
    }
    
	for _, table := range tables {

		actual:= stringShift(table.a, table.b);

		if actual != table.exp{
			t.Errorf("%s, %d wrong string shift, actual %s should be %s", table.a, table.b, actual, table.exp)
         }
	}
}