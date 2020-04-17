package main;

func checkValidString(s string) bool {

	left := 0;
	wildcard:=0;
	for _, char := range s {

		if char == '('{
			left++;
		}else if char == '*' {
			wildcard++;
		}else if char == ')' {
			if left <= 0{
				if wildcard > 0 {
					wildcard--;
					continue;
				}
				return false;
			}
			left--;
		}
	}

	wildcard =0;
	right:= 0;
	for i := len(s) -1; i >=0; i-- {
		char := s[i]
		if char == ')'{
			right++;
		}else if char == '*' {
			wildcard++;
		}else if char == '(' {
			if right <= 0{
				if wildcard > 0 {
					wildcard--;
					continue;
				}
				return false;
			}
			right--;
		}
	}
	return true;
}
