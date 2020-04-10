package main

import "testing"

func TestGetMinInStack(t *testing.T) {
	tables := []struct {
		a []int
		exp int
	}{
		{ []int {1,2,3}, 1},
		{ []int {2,3,4}, 2},
    }
    
	for _, table := range tables {
		obj := Constructor();
		
		param_3 := obj.Top();
		param_4 := obj.GetMin();

		for _, n := range table.a {
			obj.Push(n);
		}
		actualMin:= obj.GetMin();

        if actualMin != table.exp {
			t.Errorf("%d, min not equal not equal %s, should be %d", table.a, actualMin, table.exp);
         }
	}
}