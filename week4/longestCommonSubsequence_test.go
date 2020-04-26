package main;

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	tables := []struct {
		a string
		b string
		exp int
	}{
		{ "ab", "d", 0},
		{ "abc", "def", 0},
		{ "ab", "zb", 1},
		{ "ab", "az", 1},
		{ "ab", "bz", 1},
		{ "ac", "az", 1},
		{ "ac", "az", 1},
		{ "abcde", "abc", 3},
    }
    
	for _, table := range tables {
		result:= longestCommonSubsequence(table.a, table.b);

        if result != table.exp {
			t.Errorf("%s longest sequence in %s, %d should be %d", table.a, table.b, result, table.exp);
         }
	}
}

