package main

import "testing"

func TestCompareStringsWithBackspace(t *testing.T) {
	tables := []struct {
		a string
		b string
	}{
		{ "AB", "AB"},
		{ "A#", "B#"},
		{ "A#C", "B#C"},
		{ "AB#B", "AC#B"},
		{ "ABB###C", "ACA###C"},
    }
    
	for _, table := range tables {
        result := backspaceCompare(table.a, table.b)
        if !result {
            t.Errorf("Strings are not equal %s, %s", table.a, table.b)
         }
	}
}

func TestCompareStringsWithBackspaceFalse(t *testing.T) {
	tables := []struct {
		a string
		b string
	}{
		{ "AB", "AC"},
		{ "A#BB", "A#C"},
		{ "bxj##tw", "bxj###tw"},
    }
    
	for _, table := range tables {
        result := backspaceCompare(table.a, table.b)
        if result {
            t.Errorf("Strings are equal but should not be %s, %s", table.a, table.b)
         }
	}
}
