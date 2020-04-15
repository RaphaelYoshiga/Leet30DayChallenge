package main

import "testing"

func TestGetMinInStack(t *testing.T) {
	tables := []struct {
		a []int
		exp []int
	}{
		{ []int {0,0}, []int { 0,0 }},
		{ []int {1,0}, []int { 0,1 }},
		{ []int {1,2,3}, []int { 6, 3, 2 }},
		{ []int {1,2,3,4}, []int { 24,12,8,6 }},
	}
    
	for _, table := range tables {

		result:=  productExceptSelf(table.a);

		for i, actual:= range result{
			if actual != table.exp[i] {
				t.Errorf("%d, wrong product except self: %d should be %d", table.a, result, table.exp);
			 }
		}

     
	}
}