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
		
		// param_3 := obj.Top();
		// param_4 := obj.GetMin();

		for _, n := range table.a {
			obj.Push(n);
		}
		actualMin:= obj.GetMin();

        if actualMin != table.exp {
			t.Errorf("%d, min not equal %d, should be %d", table.a, actualMin, table.exp);
         }
	}
}

func TestPopTopMinStack(t *testing.T) {
	tables := []struct {
		a []int
		exp int
	}{
		{ []int {1,2,3}, 2},
		{ []int {2,3,4}, 3},
    }
    
	for _, table := range tables {
		obj := Constructor();
		
		// param_3 := obj.Top();
		// param_4 := obj.GetMin();

		for _, n := range table.a {
			obj.Push(n);
		}
		obj.Pop();
		top:= obj.Top();

        if top != table.exp {
			t.Errorf("%d, top not equal %d, should be %d", table.a, top, table.exp);
         }
	}
}


func TestPopMinStack(t *testing.T) {
	tables := []struct {
		a []int
		exp int
	}{
		{ []int {3,2,1}, 2},
		{ []int {4,3,1}, 3},
    }
    
	for _, table := range tables {
		obj := Constructor();
		
		for _, n := range table.a {
			obj.Push(n);
		}
		obj.Pop();

		actualMin:= obj.GetMin();

        if actualMin != table.exp {
			t.Errorf("%d, min not equal %d, should be %d", table.a, actualMin, table.exp);
         }
	}
}