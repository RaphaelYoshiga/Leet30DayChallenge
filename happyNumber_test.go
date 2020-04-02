package main

import "testing"

func TestSumManyHappy(t *testing.T) {
	tables := []struct {
		x int
	}{
		{19},
        {1},
		{13},        
    }
    
	for _, table := range tables {

        print(table.x);
        
        result := isHappy(table.x)
        print("\n")
        print(result)
        if !(result) {
            t.Errorf("%d is a happy number! but it was not", table.x)
         }
	}
}


func TestSumManyNonHappy(t *testing.T) {
	tables := []struct {
		x int
	}{
		{2},
		{4},
    }
    
	for _, table := range tables {

        result := isHappy(table.x)
        if result {
           t.Errorf("%d is not a happy number! but it was", table.x)
        }
	}
}