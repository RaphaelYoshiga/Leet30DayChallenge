package main

import "testing"

func TestValidString(t *testing.T) {
	tables := []struct {
		a string
	}{
		{ "()" },
		{ "(*)" },
		{ "(*)(*)(*)(*)" },
		{ "(*))" },
		{ "(*()" },
		{ "((*)" },
		{ "((*)" },
		{ "(((******))" },
		{ "((()))()(())(*()()())**(())()()()()((*()*))((*()*)" },
		{ "(())((())()()(*)" },

		 
	}
    
	for _, table := range tables {
		result:=  checkValidString(table.a);
		if !result {
			t.Errorf("False, string is not valid: %s; but it should be", table.a);
		 }
	}
}

func TestInvalidString(t *testing.T) {
	tables := []struct {
		a string
	}{
		{ ")(" },
		{ ")()" },
		{ "())(" },
		{ ")())" },
		{ "(*()(())())())()()((()())((()))(*"},
		{ "((*)(*))((*"},				
		{ "(())((())()()(*)(*()(())())())()()((()())((()))(*"},
	}
    
	for _, table := range tables {
		result:=  checkValidString(table.a);
		if result {
			t.Errorf("True, string is valid: %s ; but should be false.", table.a);
		 }
	}
}